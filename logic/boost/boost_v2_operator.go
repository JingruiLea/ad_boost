package boost

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/envs"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/lark"
	"github.com/JingruiLea/ad_boost/logic/ad"
	"github.com/JingruiLea/ad_boost/logic/ad_report"
	"github.com/JingruiLea/ad_boost/logic/boost/operator"
	"github.com/JingruiLea/ad_boost/logic/live_report"
	"github.com/JingruiLea/ad_boost/model/bo"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/jinzhu/now"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"sort"
	"strings"
	"time"
)

var ops = []*BoostOperator{
	{
		AdGroupID: 1788144214349987,
		AccountID: 1784698853978186,
		AwemeID:   2691211639665967,
		RealTimeParam: RealTimeParam{
			Ads:           nil,
			ReportsMinMap: make(map[int64][]*ad_report.AdReport),
			CostMinMap:    make(map[int64]float64),
		},
		BoostParam: BoostParam{
			CheckPeriodMin: 5,
			TotalBudget:    2000,
			MinRoi:         2.5,
			BudgetTimeMap: map[int64]int64{
				120: 3000,
			},
			RoiTimeMap: map[int64]float64{
				120: 6,
			},
			CostThreshold:    200,
			MinCostThreshold: 2,
			CanStopMin:       5,
			RecoverRoi:       5,
			WarningAdCanKeep: 10,
		},
		Stopped: true,
	},
}

type BoostOperator struct {
	AdGroupID  int64      `json:"ad_group_id"` // 广告组ID // 1788144214349987
	AccountID  int64      `json:"account_id"`  // 广告主ID // 1784698853978186
	AwemeID    int64      `json:"aweme_id"`    // 抖音号ID // 2691211639665967
	AdIDs      []int64    `json:"-"`           //监听的广告ID
	BoostParam `json:"-"` //超参数可以改和更新, 不要存
	RealTimeParam
	Stopped bool `json:"stopped"` //是否已经停止
	Op      operator.AdOperator
	dryRun  bool
}

// 实时参数
type RealTimeParam struct {
	Ads           []*bo.Ad
	ReportsMinMap map[int64][]*ad_report.AdReport
	RoomStartTime time.Time //直播间开播时间
	CostMinMap    map[int64]float64
	StartMinCount int64
	WarningAdMap  map[int64]int64 //危险计划, adID vs 开始时间
	LastCost      float64         //上一次的消耗
}

// 超参数
type BoostParam struct {
	TotalBudget      int64             `json:"total_budget"`        // 总预算,2小时总预算
	MinRoi           float64           `json:"min_roi"`             // 最小ROI
	BudgetTimeMap    map[int64]int64   `json:"budget_time_map"`     // 预算时间段, 分钟vs总cost,单位元
	RoiTimeMap       map[int64]float64 `json:"roi_time_map"`        // roi时间段, 分钟vsRoi
	CheckPeriodMin   int64             `json:"check_period_min"`    // 检查周期, 分钟
	CostThreshold    float64           `json:"cost_threshold"`      // 消耗阈值, 消耗超过这个值,就降价
	MinCostThreshold float64           `json:"min_cost_threshold"`  // 消耗阈值, 消耗小于这个值,就涨价
	CanStopMin       int64             `json:"can_stop_min"`        // 可以停止的分钟数
	RecoverRoi       float64           `json:"recover_roi"`         // 恢复ROI
	WarningAdCanKeep int64             `json:"warning_ad_can_keep"` // 危险计划可以保留的分钟数
}

var operatorActions = []string{
	"start_ad",          //开启计划
	"stop_ad",           //暂停计划
	"add_bid_to",        //加价
	"sub_bid_to",        //降价
	"add_budget_to",     //加预算
	"sub_budget_to",     //减预算
	"create_ad",         //添加广告
	"delete_ad",         //删除广告
	"copy_ad",           //复制广告
	"add_roi_target_to", //增加ROI目标
	"sub_roi_target_to", //减少ROI目标
}

//创建广告的四大步骤
//1. 预算
//2. 目标: bid price, roi target
//3. audience: 圈选人群
//4. creative: 创意 optional

func InitOperators() {
	ctx := context.Background()
	if envs.IsDev() {
		logs.CtxInfof(ctx, "dev mode, skip InitOperators")
		return
	}
	lark.RegisterTextHandler(context.Background(), func(msg string) string {
		switch {
		case strings.HasPrefix(msg, "start_ad"):
			for _, op := range ops {
				err := op.Start(ctx)
				if err != nil {
					logs.CtxErrorf(ctx, "op.Start error: %v", err)
					continue
				}
			}
			return "success start"
		case strings.HasPrefix(msg, "stop_ad"):
			utils.SafeGo(ctx, func() {
				for _, op := range ops {
					err := op.Stop(ctx)
					if err != nil {
						logs.CtxErrorf(ctx, "op.Stop error: %v", err)
						continue
					}
				}
			})
			return "success stop"
		}
		return "OK"
	})
}
func (op *BoostOperator) Save(ctx context.Context) error {
	rc := redis_dal.GetRedisClient()
	key := fmt.Sprintf("boost:operator:%d:%d", op.AccountID, op.AwemeID)
	val := utils.GetJsonStr(op)
	err := rc.Set(ctx, key, val, time.Hour*24*7).Err()
	if err != nil {
		logs.CtxErrorf(ctx, "rc.Set error: %v", err)
		return err
	}
	logs.CtxInfof(ctx, "save to redis key:%s, val:%s", key, val)
	return nil
}

func (op *BoostOperator) Load(ctx context.Context) error {
	rc := redis_dal.GetRedisClient()
	key := fmt.Sprintf("boost:operator:%d:%d", op.AccountID, op.AwemeID)
	val, err := rc.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		logs.CtxErrorf(ctx, "rc.Get error: %v", err)
		return err
	}
	d := jsoniter.NewDecoder(bytes.NewBuffer([]byte(val)))
	d.UseNumber()
	err = d.Decode(op)
	if err != nil {
		logs.CtxErrorf(ctx, "d.Decode error: %v", err)
		return err
	}
	return nil
}

func (op *BoostOperator) Start(ctx context.Context) error {
	//1. 从redis中加载
	err := op.Load(ctx)
	if err != nil {
		logs.CtxErrorf(ctx, "Load error: %v", err)
		return err
	}
	//2. 启动循环
	utils.SafeGo(ctx, func() {
		op.Loop(ctx)
	})
	return nil
}

func (op *BoostOperator) Stop(ctx context.Context) error {
	op.Stopped = true
	time.Sleep(time.Minute) //休息一分钟 确保停止
	return nil
}

func (op *BoostOperator) Restart(ctx context.Context) error {
	op.Stopped = false
	return nil
}

func (op *BoostOperator) Loop(ctx context.Context) {
	larkAndLog(ctx, "开始循环")
	//每分钟检测一次
	for range time.Tick(time.Minute) {
		err := op.Once(ctx)
		if err != nil {
			logs.CtxErrorf(ctx, "Once error: %v", err)
			continue
		}
		err = op.Save(ctx)
		if err != nil {
			logs.CtxErrorf(ctx, "Save error: %v", err)
			continue
		}
		if op.Stopped {
			logs.CtxInfof(ctx, "stopped, break")
			break
		}
	}
}

func (op *BoostOperator) Once(ctx context.Context) error {
	nowList, err := live_report.GetNowLiveRoomList(ctx, op.AccountID, op.AwemeID)
	if err != nil {
		logs.CtxErrorf(ctx, "GetNowLiveRoomList error: %v", err)
		return err
	}
	if len(nowList.RoomList) == 0 {
		logs.CtxInfof(ctx, "no live room")
		larkAndLog(ctx, "no live room")
		return nil
	}
	if len(nowList.RoomList) > 1 {
		logs.CtxInfof(ctx, "more than one live room")
		larkAndLog(ctx, "more than one live room")
		return nil
	}
	err = op.LiveReport(ctx)
	if err != nil {
		logs.CtxErrorf(ctx, "LiveReport error: %v", err)
	}
	room := nowList.RoomList[0]
	roomStartTime, err := time.ParseInLocation("2006-01-02 15:04:05", room.StartTime, time.Local)
	if err != nil {
		logs.CtxErrorf(ctx, "time.Parse error: %v", err)
		return err
	}
	op.RoomStartTime = roomStartTime
	logs.CtxInfof(ctx, "roomStartTime: %s", roomStartTime.Format("2006-01-02 15:04:05"))
	startMins := int64(time.Since(roomStartTime).Minutes())
	op.StartMinCount = startMins
	logs.CtxInfof(ctx, "startMins: %d", startMins)
	larkAndLog(ctx, "开始第n分钟的检测: n:%d", startMins)
	//看看自己组里的广告, 一般不会超过100个
	ads, err := op.GetAllAd(ctx, "")
	if err != nil {
		logs.CtxErrorf(ctx, "GetAllAd error: %v", err)
		return err
	}
	if ads == nil || len(ads) == 0 {
		logs.CtxInfof(ctx, "no ad")
		larkAndLog(ctx, "no ad")
		return nil
	}
	adIDs := make([]int64, 0, len(ads))
	for _, aditem := range ads {
		adIDs = append(adIDs, aditem.AdID)
	}
	adMap := make(map[int64]*bo.Ad, len(ads))
	for _, aditem := range ads {
		adMap[aditem.AdID] = aditem
	}
	reports, err := ad_report.MGetCommonAdDailyReport(ctx, op.AccountID, adIDs)
	if err != nil {
		logs.CtxErrorf(ctx, "ad_report.MGetCommonAdReport failed, err:%v", err)
		larkAndLog(ctx, fmt.Sprintf("ad_report.MGetCommonAdReport failed, err:%v", err))
		return err
	}
	if reports == nil || len(reports) == 0 {
		logs.CtxInfof(ctx, "no report")
		larkAndLog(ctx, "no report")
	}

	if op.ReportsMinMap == nil {
		op.ReportsMinMap = map[int64][]*ad_report.AdReport{
			startMins: reports,
		}
	} else {
		op.ReportsMinMap[startMins] = reports
	}

	totalPayAmount := float64(0)
	totalCost := float64(0)
	totalRoi := float64(0)

	sort.Slice(reports, func(i, j int) bool {
		return reports[i].StatCost > reports[j].StatCost
	})

	for _, report := range reports {
		totalPayAmount += report.PayOrderAmount
		totalCost += report.StatCost
	}
	totalRoi = utils.RoundFloat(totalPayAmount/totalCost, 2)
	if op.CostMinMap == nil {
		op.CostMinMap = map[int64]float64{
			startMins: totalCost,
		}
	} else {
		op.CostMinMap[startMins] = totalCost
	}

	larkAndLog(ctx, "总消耗:%.2f, 总成交金额:%.2f, 总ROI:%.2f", totalCost, totalPayAmount, totalRoi)
	avgCost := totalCost/float64(len(reports)) + 0.01
	reportStr := ""
	mainStr := ""
	for _, report := range reports {
		ad := adMap[report.AdID]
		if report.StatCost > avgCost { //大于均值的是主要计划
			mainStr += fmt.Sprintf("主要计划:ID:%d, 消耗:%.2f, ROI:%.2f, 状态:%s\n", report.AdID, report.StatCost, report.PrepayAndPayOrderRoi, ad.Status)
		}
		reportStr += fmt.Sprintf("计划ID:%d, 出价:%.2f, 消耗:%.2f, ROI:%.2f, 成单:%d, 状态:%s\n",
			report.AdID, ad.DeliverySetting.CPABid,
			report.StatCost, report.PrepayAndPayOrderRoi,
			report.PayOrderCount, ad.Status)
	}
	if mainStr != "" {
		larkAndLog(ctx, mainStr)
	}
	larkAndLog(ctx, reportStr)
	if startMins%op.CheckPeriodMin == 0 {
		pCost := totalCost - op.LastCost
		if pCost > op.CostThreshold && op.LastCost > 0 {
			//整体降价10%
			larkAndLog(ctx, "最近%d分钟消耗:%.2f, 超过阈值:%.2f, 整体降价10%%", op.CheckPeriodMin, pCost, op.CostThreshold)
			for _, report := range reports {
				item := adMap[report.AdID]
				err = ad.UpdateAdBid(ctx, op.AccountID, []*ad.Bid{
					{
						AdId: report.AdID,
						Bid:  utils.RoundFloat(item.DeliverySetting.CPABid*0.9, 2),
					},
				})
				if err != nil {
					logs.CtxErrorf(ctx, "ad.UpdateAdBudget failed, err:%v", err)
					lark.SendRoomMessage(ctx, fmt.Sprintf("ad.UpdateAdBudget failed, err:%v", err))
					continue
				}
			}
		}
		if pCost < op.MinCostThreshold && op.LastCost > 0 {
			//整体提价10%
			larkAndLog(ctx, "最近%d分钟消耗:%.2f, 小于阈值:%.2f, 整体提价10%%", op.CheckPeriodMin, pCost, op.MinCostThreshold)
			for _, report := range reports {
				item := adMap[report.AdID]
				err = ad.UpdateAdBid(ctx, op.AccountID, []*ad.Bid{
					{
						AdId: report.AdID,
						Bid:  utils.RoundFloat(item.DeliverySetting.CPABid*1.1, 2),
					},
				})
				if err != nil {
					logs.CtxErrorf(ctx, "ad.UpdateAdBudget failed, err:%v", err)
					lark.SendRoomMessage(ctx, fmt.Sprintf("ad.UpdateAdBudget failed, err:%v", err))
					continue
				}
			}
		}
		op.LastCost = totalCost
		for _, report := range reports {
			if report.StatCost < 100.0 && report.PrepayAndPayOrderRoi == 0 { //消耗小于100块而且没有成单,可以先忽略
				larkAndLog(ctx, "计划ID:%d, 消耗:%.2f, ROI:%.2f, 成单:%d, 状态:%s, 消耗太低,忽略",
					report.AdID, report.StatCost, report.PrepayAndPayOrderRoi, report.PayOrderCount, adMap[report.AdID].Status)
				continue
			}
			item := adMap[report.AdID]
			if report.PrepayAndPayOrderRoi <= op.MinRoi {
				if op.StartMinCount <= op.CanStopMin {
					larkAndLog(ctx, "还没到可以停止的时间, 忽略")
					continue
				}
				//如果已经是危险的了
				startMin, ok := op.WarningAdMap[item.AdID]
				if ok && op.StartMinCount-startMin > op.WarningAdCanKeep {
					err = ad.UpdateAdStatus(ctx, &ad.UpdateAdStatusReq{
						AdvertiserID: op.AccountID,
						AdIDs:        []int64{item.AdID},
						OptStatus:    ttypes.OptStatusDisable,
					})
					if err != nil {
						logs.CtxErrorf(ctx, "ad.UpdateAdStatus failed, err:%v", err)
						lark.SendRoomMessage(ctx, fmt.Sprintf("ad.UpdateAdStatus failed, err:%v", err))
						continue
					}
					larkAndLog(ctx, "暂停计划:ID:%d,因为ROI太低了:%.2f", item.AdID, report.PrepayAndPayOrderRoi)
				}
				//设置为危险计划
				if op.WarningAdMap == nil {
					op.WarningAdMap = map[int64]int64{
						item.AdID: op.StartMinCount,
					}
				} else {
					op.WarningAdMap[item.AdID] = op.StartMinCount
				}
				larkAndLog(ctx, "设置为危险计划:ID:%d", item.AdID)
			}
			//roi恢复逻辑
			for adID, m := range warningAdMap {
				if m.AdID == item.AdID {
					//如果已经是危险的了
					if report.PrepayAndPayOrderRoi > op.RecoverRoi {
						delete(warningAdMap, adID)
						larkAndLog(ctx, "重启危险计划:ID:%d,因为ROI恢复了:%.2f", item.AdID, report.PrepayAndPayOrderRoi)
						err = ad.UpdateAdStatus(ctx, &ad.UpdateAdStatusReq{
							AdvertiserID: op.AccountID,
							AdIDs:        []int64{adID},
							OptStatus:    ttypes.OptStatusEnable,
						})
						if err != nil {
							logs.CtxErrorf(ctx, "ad.UpdateAdStatus failed, err:%v", err)
							lark.SendRoomMessage(ctx, fmt.Sprintf("ad.UpdateAdStatus failed, err:%v", err))
							continue
						}
					}
				}
			}
		}
	}
	return nil
}

func (op *BoostOperator) NewAd(ctx context.Context, audienceID int64) error {
	//高-0107-三笛-自定义-控成本-7ROI2.4-女4149排负向-直投
	//（操作人，保持【高】不变）-（日期，用四位数字表示）-（直播间编号+主播昵称，我们先用客户第三直播间，笛笛这个主播来测试）-（推广方式）-（投放方式）-（优化目标+出价，例如roi2=roi出价，出价为2, 7roi3=7日roi，出价为3）-定向包-素材
	adIns := bo.NewLiveCommonAd("", op.AwemeID, op.AccountID, op.AdGroupID).WithBudget(300).WithBid(500)
	adIns.Audience = bo.NewAudience().WithOrientationID(audienceID)
	adIns.Name = "高-0107-三笛-自定义-控成本-7ROI2.4-女4149排负向-直投"
	_, err := ad.CreateAd(ctx, adIns)
	if err != nil {
		logs.CtxErrorf(ctx, "CreateAd error: %v", err)
		return err
	}
	return nil
}

func (op *BoostOperator) GetAllAd(ctx context.Context, status ttypes.AdStatus) ([]*bo.Ad, error) {
	ads, err := ad.GetAdList(ctx, &ad.GetAdListReq{
		AdvertiserId: op.AccountID,
		Filtering: &ad.Filter{
			CampaignScene:  []ad.CampaignSceneFilter{ad.CampaignSceneFilterDailySale},
			MarketingGoal:  ttypes.MarketingGoalLivePromGoods,
			MarketingScene: ad.MarketingSceneFilterFeed,
			Status:         status,
			CampaignId:     op.AdGroupID,
			AwemeId:        op.AwemeID,
		},
		Page:     1,
		PageSize: 50,
	})
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdList error: %v", err)
		return nil, err
	}
	op.Ads = ads.List
	return op.Ads, nil
}

func (op *BoostOperator) LiveReport(ctx context.Context) error {
	//直播间数据
	liveDatas, err := live_report.GetLiveRoomList(ctx, &live_report.GetLiveRoomListReq{
		AdvertiserID: op.AccountID,
		AwemeID:      op.AwemeID,
		DateTime:     time.Now().Format("2006-01-02"),
		RoomStatus:   live_report.RoomStatusLiving,
		AdStatus:     live_report.AdStatusAll,
		Fields:       live_report.RoomMetricsFieldStatCost.All(),
		Page:         1,
		PageSize:     10,
	})
	if err != nil {
		logs.CtxErrorf(ctx, "live_report.GetLiveRoomList failed, err:%v", err)
		return err
	}
	if len(liveDatas.RoomList) == 0 {
		//说明是昨天开播的今天的直播间
		liveDatas, err = live_report.GetLiveRoomList(ctx, &live_report.GetLiveRoomListReq{
			AdvertiserID: op.AccountID,
			AwemeID:      op.AwemeID,
			DateTime:     now.BeginningOfDay().Add(time.Minute * -1).Format("2006-01-02"),
			RoomStatus:   live_report.RoomStatusLiving,
			AdStatus:     live_report.AdStatusAll,
			Fields:       live_report.RoomMetricsFieldStatCost.All(),
			Page:         1,
			PageSize:     10,
		})
		if err != nil {
			logs.CtxErrorf(ctx, "live_report.GetLiveRoomList failed, err:%v", err)
			return err
		}
		if len(liveDatas.RoomList) == 0 {
			logs.CtxErrorf(ctx, "报错: 今天和昨天都没有直播间数据")
			return nil
		}
	}
	var roomRoi, roomCost, roomGpm, roomPayAmount, adGpm, adPayAmount, adRoi float64
	if len(liveDatas.RoomList) > 0 {
		room := liveDatas.RoomList[0]
		roomRoi = room.LivePayOrderGmvRoi                                     //整体ROI
		roomCost = utils.RoundFloat(room.StatCost/100000, 2)                  //整体消耗
		roomGpm = utils.RoundFloat(room.TotalLivePayOrderGpm/100000, 2)       //整体千次观看成交
		roomPayAmount = utils.RoundFloat(room.LivePayOrderGmvAlias/100000, 2) //整体成交金额

		adGpm = utils.RoundFloat(room.LubanLivePayOrderGpm/100000, 2)       //广告千次观看成交
		adPayAmount = utils.RoundFloat(room.LubanLivePayOrderGmv/100000, 2) //广告成交金额
		adRoi = room.AdLivePrepayAndPayOrderGmvRoi                          //广告ROI
	}
	larkAndLog(ctx, "直播间数据: 广告ROI:%.2f, 消耗:%.2f\n 整体ROI:%.2f, 整体GPM:%.2f, 整体成交金额:%.2f\n广告数据: GPM:%.2f, 成交金额:%.2f", adRoi, roomCost, roomRoi, roomGpm, roomPayAmount, adGpm, adPayAmount)
	return nil
}
