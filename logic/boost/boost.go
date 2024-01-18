package boost

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/ad_dal"
	"github.com/JingruiLea/ad_boost/lark"
	"github.com/JingruiLea/ad_boost/logic/ad"
	"github.com/JingruiLea/ad_boost/logic/ad_report"
	"github.com/JingruiLea/ad_boost/logic/boost/sync"
	"github.com/JingruiLea/ad_boost/logic/live_report"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/jinzhu/now"
	"sort"
	"time"
)

var boostOn bool

var fakeRun bool

var boostStart *time.Time

var cost float64

var intervalSeconds int64 = 1

// warning_ad_map
var warningAdMap = make(map[int64]*model.Ad)

type CmdFilter func(cmd string) bool

func Init() {
	BoostV2MonitorStart(context.Background())
	//lark.RegisterTextHandler(context.Background(), func(msg string) string {
	//	logs.CtxInfof(context.Background(), "收到消息:%s", msg)
	//	switch {
	//	case strings.HasPrefix(msg, "boost init"):
	//		accountID := utils.Str2I64(strings.Split(msg, " ")[2])
	//		if accountID == 0 {
	//			return "boost init failed,参数不对"
	//		}
	//		err := BoostInit(context.Background(), accountID)
	//		if err != nil {
	//			return fmt.Sprintf("boost init failed, err:%v", err)
	//		}
	//		return "boost init done"
	//	case strings.HasPrefix(msg, "boost start"):
	//		var err error
	//		adArr := strings.Split(msg, " ")
	//		if len(adArr) == 4 {
	//			startTime, err = time.ParseInLocation("2006-01-0215:04", adArr[2]+adArr[3], time.Local)
	//			if err != nil {
	//				larkAndLog(context.Background(), "time.Parse failed, err:%v", err)
	//			}
	//		}
	//		accountID := utils.Str2I64(adArr[2])
	//		if accountID == 0 {
	//			return "boost start failed,参数不对"
	//		}
	//		BoostStart(context.Background(), 1748031128935424, startTime)
	//		return "boost start"
	//	case strings.HasPrefix(msg, "boost stop"):
	//		boostOn = false
	//		tickCount = 0
	//		return "boost stop"
	//	case strings.HasPrefix(msg, "boost status"):
	//		return fmt.Sprintf("boost status: %v", boostOn)
	//	case strings.HasPrefix(msg, "boost interval"):
	//		adArr := strings.Split(msg, " ")
	//		if len(adArr) != 3 {
	//			return "boost interval failed,参数不对"
	//		}
	//		intervalSecond := utils.Str2I64(adArr[2])
	//		if intervalSecond == 0 {
	//			return "boost interval failed,参数不对"
	//		}
	//		intervalSeconds = intervalSecond
	//		return fmt.Sprintf("boost interval:%d", intervalSeconds)
	//	case strings.HasPrefix(msg, "boost fake run"):
	//		fakeRun = true
	//		return "boost fake run"
	//	case strings.HasPrefix(msg, "boost fake stop"):
	//		fakeRun = false
	//		//实时报表
	//	case strings.HasPrefix(msg, "boost report start"):
	//		AlarmOpen = true
	//		return "boost report"
	//	case strings.HasPrefix(msg, "boost report stop"):
	//		AlarmOpen = false
	//	case strings.HasPrefix(msg, "更新出价"):
	//		//更新出价 ID 出价
	//		//更新出价 123 0.5
	//		adArr := strings.Split(msg, " ")
	//		if len(adArr) != 3 {
	//			return "更新出价失败,参数不对"
	//		}
	//		adID := utils.Str2I64(adArr[1])
	//		if adID == 0 {
	//			return "更新出价失败,参数不对"
	//		}
	//		bid := utils.Str2Float64(adArr[2])
	//		if bid == 0 {
	//			return "更新出价失败,参数不对"
	//		}
	//		err := ad.UpdateAdBid(context.Background(), 1748031128935424, []*ad.Bid{
	//			{
	//				AdId: adID,
	//				Bid:  bid,
	//			},
	//		})
	//		if err != nil {
	//			return fmt.Sprintf("更新出价失败, err:%v", err)
	//		}
	//
	//	case strings.HasPrefix(msg, "停止"):
	//		//停止 ID
	//		adArr := strings.Split(msg, " ")
	//		if len(adArr) != 2 {
	//			return "停止失败,参数不对"
	//		}
	//		adID := utils.Str2I64(adArr[1])
	//		if adID == 0 {
	//			return "停止失败,参数不对"
	//		}
	//		err := ad.UpdateAdStatus(context.Background(), &ad.UpdateAdStatusReq{
	//			AdvertiserID: 1748031128935424,
	//			AdIDs:        []int64{adID},
	//			OptStatus:    ttypes.OptStatusDisable,
	//		})
	//		if err != nil {
	//			return fmt.Sprintf("停止失败, err:%v", err)
	//		}
	//
	//	case strings.HasPrefix(msg, "启动"):
	//		//启动 ID
	//		adArr := strings.Split(msg, " ")
	//		if len(adArr) != 2 {
	//			return "启动失败,参数不对"
	//		}
	//		adID := utils.Str2I64(adArr[1])
	//		if adID == 0 {
	//			return "启动失败,参数不对"
	//		}
	//		err := ad.UpdateAdStatus(context.Background(), &ad.UpdateAdStatusReq{
	//			AdvertiserID: 1748031128935424,
	//			AdIDs:        []int64{adID},
	//			OptStatus:    ttypes.OptStatusEnable,
	//		})
	//		if err != nil {
	//			return fmt.Sprintf("启动失败, err:%v", err)
	//		}
	//	}
	//	return ""
	//})
}

func BoostInit(ctx context.Context, accountID int64) error {
	//同步广告
	err := sync.SyncAds(ctx, accountID)
	if err != nil {
		logs.CtxErrorf(ctx, "sync.SyncAds failed, err:%v", err)
		return err
	}
	//获取所有广告
	ads, err := ad_dal.MGetAdsByAdvertiserID(ctx, accountID)
	if err != nil {
		logs.CtxErrorf(ctx, "ad_dal.MGetAdsByAdvertiserID failed, err:%v", err)
		return err
	}
	for _, m := range ads {
		if m.Status == ttypes.AdStatusDeliveryOk {
			larkAndLog(ctx, "广告正在投放,我不敢操作.")
			return nil
		}
	}
	//把所有广告预算改到DEFAULT, 防止爆炸
	budgets := make([]*ad.Budget, 0, len(ads))
	for _, item := range ads {
		budgets = append(budgets, &ad.Budget{
			AdID:   item.AdID,
			Budget: DefaultBudget,
		})
	}
	err = ad.MUpdateAdBudgets(ctx, accountID, budgets)
	if err != nil {
		logs.CtxErrorf(ctx, "ad.UpdateAdBudget failed, err:%v", err)
		return err
	}
	larkAndLog(ctx, "更新广告预算到默认值成功")
	//更新ROI到默认值
	var adIDs []int64
	for _, item := range ads {
		adIDs = append(adIDs, item.AdID)
		if item.DeliverySetting.DeepExternalAction == ttypes.DeepExternalActionAdConvertTypeLivePayRoi {
			err := ad.UpdateAdRoiGoal(ctx, accountID, []*ad.RoiGoal{
				{
					AdId:    item.AdID,
					RoiGoal: DefaultRoi,
				},
			})
			if err != nil {
				logs.CtxErrorf(ctx, "ad.UpdateAdRoiGoal failed, err:%v", err)
				return err
			}
			larkAndLog(ctx, "更新ROI到默认值成功:ID:%d, ROI:%.2f", item.AdID, DefaultRoi)
		} else if item.DeliverySetting.ExternalAction == ttypes.ExternalActionAdConvertTypeLiveSuccessorderPay {
			err := ad.UpdateAdBid(ctx, accountID, []*ad.Bid{
				{
					AdId: item.AdID,
					Bid:  DefaultBid,
				},
			})
			if err != nil {
				logs.CtxErrorf(ctx, "ad.UpdateAdBid failed, err:%v", err)
				return err
			}
			larkAndLog(ctx, "更新出价到默认值成功:ID:%d, 出价:%.2f", item.AdID, DefaultBid)
		} else if item.DeliverySetting.ExternalAction == ttypes.ExternalActionAdConvertTypeNewFollowAction {
			larkAndLog(ctx, "找到关注广告:ID:%d", item.AdID)
			err := ad.UpdateAdBid(ctx, accountID, []*ad.Bid{
				{
					AdId: item.AdID,
					Bid:  DefaultFollowUpBid,
				},
			})
			if err != nil {
				logs.CtxErrorf(ctx, "ad.UpdateAdBid failed, err:%v", err)
				return err
			}
			larkAndLog(ctx, "更新关注出价到默认值成功:ID:%d, 出价:%.2f", item.AdID, DefaultFollowUpBid)
		}
	}
	larkAndLog(ctx, "更新ROI和出价到默认值成功")
	//启动所有广告
	err = ad.MUpdateAdStatus(ctx, accountID, adIDs, ttypes.OptStatusEnable)
	if err != nil {
		logs.CtxErrorf(ctx, "ad.MUpdateAdStatus failed, err:%v", err)
		return err
	}
	larkAndLog(ctx, "启动所有广告成功, 开始启动定时监测")
	//同步广告
	err = sync.SyncAds(ctx, accountID)
	if err != nil {
		logs.CtxErrorf(ctx, "sync.SyncAds failed, err:%v", err)
		return err
	}
	return nil
}

func BoostStart(ctx context.Context, accountID int64, t time.Time) {
	larkAndLog(ctx, "开始启动定时监测,V2")
	if !t.IsZero() {
		larkAndLog(ctx, "指定时间:%s", t.Format("2006-01-02 15:04"))
		boostStart = &t
	}
	BoostLoop(ctx, accountID, t)
}

var startTime time.Time
var tickCount int

func BoostLoop(ctx context.Context, accountID int64, t time.Time) {
	boostOn = true
	BoostOnce(ctx, accountID)
	if tickCount != 0 {
		return
	}
	utils.SafeGo(ctx, func() {
		for range time.Tick(time.Duration(int64(time.Minute) * intervalSeconds)) {
			if !boostOn {
				larkAndLog(ctx, "定时监测关闭")
				tickCount = 0
				return
			}
			tickCount++
			BoostOnce(ctx, accountID)
		}
	})
}

func BoostOnce(ctx context.Context, accountID int64) {
	//正在投放
	totalAds, err := ad.GetAdListByStatus(ctx, accountID, ttypes.AdStatusDeliveryOk, 0, nil)
	if err != nil {
		logs.CtxErrorf(ctx, "ad.GetAdListByStatus failed, err:%v", err)
		lark.SendRoomMessage(ctx, fmt.Sprintf("ad.GetAdListByStatus failed, err:%v", err))
		return
	}

	if len(totalAds) == 0 {
		boostStart = nil
		return
	}
	//暂停的计划
	respPause, err := ad.GetAdListByStatus(ctx, accountID, ttypes.AdStatusDisable, 0, nil)
	if err != nil {
		logs.CtxErrorf(ctx, "ad.GetAdListByStatus failed, err:%v", err)
		lark.SendRoomMessage(ctx, fmt.Sprintf("ad.GetAdListByStatus failed, err:%v", err))
		return
	}
	var pausedIDs []int64
	larkAndLog(ctx, "找到%d个正在投放的计划, %d个暂停的计划", len(totalAds), len(respPause))
	if len(respPause) > 0 {
		totalAds = append(totalAds, respPause...)
	}
	if boostStart == nil {
		now := time.Now()
		boostStart = &now
		for _, m := range respPause {
			pausedIDs = append(pausedIDs, m.AdID)
		}
		if len(pausedIDs) != 0 {
			//启动所有暂停的计划
			err = ad.UpdateAdStatus(ctx, &ad.UpdateAdStatusReq{
				AdvertiserID: accountID,
				AdIDs:        pausedIDs,
				OptStatus:    ttypes.OptStatusEnable,
			})
			if err != nil {
				logs.CtxErrorf(ctx, "ad.UpdateAdStatus failed, err:%v", err)
				lark.SendRoomMessage(ctx, fmt.Sprintf("ad.UpdateAdStatus failed, err:%v", err))
				return
			}
		}
	}
	adIDs := make([]int64, 0, len(totalAds))
	for _, item := range totalAds {
		adIDs = append(adIDs, item.AdID)
	}

	adMap, err := ad_dal.MGetAdByAdIDs(ctx, adIDs)
	if err != nil {
		logs.CtxErrorf(ctx, "ad_dal.MGetAdByAdIDs failed, err:%v", err)
		lark.SendRoomMessage(ctx, fmt.Sprintf("ad_dal.MGetAdByAdIDs failed, err:%v", err))
		return
	}

	reports, err := ad_report.MGetCommonAdDailyReport(ctx, accountID, adIDs)
	if err != nil {
		logs.CtxErrorf(ctx, "ad_report.MGetCommonAdReport failed, err:%v", err)
		lark.SendRoomMessage(ctx, fmt.Sprintf("ad_report.MGetCommonAdReport failed, err:%v", err))
		return
	}
	adReportItems := make([]*model.AdReportItem, 0, len(reports))
	for _, item := range reports {
		adItem, ok := adMap[item.AdID]
		if !ok {
			logs.CtxErrorf(ctx, "!ok, adID:%d", item.AdID)
			lark.SendRoomMessage(ctx, fmt.Sprintf("!ok, adID:%d", item.AdID))
			continue
		}
		adReportItems = append(adReportItems, item.ToModel(adItem.DeliverySetting.CPABid, adItem.DeliverySetting.ROIGoal, "", 0))
	}
	err = ad_dal.CreateAdReportItem(ctx, adReportItems)
	if err != nil {
		logs.CtxErrorf(ctx, "ad_dal.CreateAdReportItem failed, err:%v", err)
		lark.SendRoomMessage(ctx, fmt.Sprintf("ad_dal.CreateAdReportItem failed, err:%v", err))
		return
	}
	lastReports, err := ad_dal.MGetLastAdReport(ctx, adIDs, boostStart.Add(time.Hour*-8), now.New(time.Now().Add(time.Hour*-8)).BeginningOfDay())
	if err != nil {
		logs.CtxErrorf(ctx, "ad_dal.MGetLastAdReport failed, err:%v", err)
		lark.SendRoomMessage(ctx, fmt.Sprintf("ad_dal.MGetLastAdReport failed, err:%v", err))
		return
	}
	lastReportMap := make(map[int64]*model.AdReportItem)
	for _, item := range lastReports {
		lastReportMap[item.AdID] = item
	}
	minutes := int(time.Since(*boostStart).Minutes())
	if minutes < 0 {
		logs.CtxErrorf(ctx, "time.Since(*boostStart).Minutes() < 0")
		lark.SendRoomMessage(ctx, fmt.Sprintf("time.Since(*boostStart).Minutes() < 0"))
		return
	}
	larkAndLog(ctx, "开始第%d分钟的监测", minutes)
	totalPayAmount := float64(0)
	totalCost := float64(0)
	totalRoi := float64(0)

	for _, report := range reports {
		lastReport, ok := lastReportMap[report.AdID]
		if !ok {
			logs.CtxInfof(ctx, "!ok, adID:%d", report.AdID)
			continue
		}
		report.PayOrderAmount = report.PayOrderAmount - (lastReport.PayOrderAmount)
		report.PayOrderCount = report.PayOrderCount - (lastReport.PayOrderCount)
		report.StatCost = report.StatCost - (lastReport.StatCost)
		report.PrepayAndPayOrderRoi = utils.RoundFloat(report.PayOrderAmount/report.StatCost, 2)
	}
	//分开
	for _, report := range reports {
		totalPayAmount += report.PayOrderAmount
		totalCost += report.StatCost
	}
	//根据消耗从大到小
	sort.Slice(reports, func(i, j int) bool {
		return reports[i].StatCost > reports[j].StatCost
	})
	totalRoi = utils.RoundFloat(totalPayAmount/totalCost, 2)
	larkAndLog(ctx, "总消耗:%.2f, 总成交金额:%.2f, 总ROI:%.2f", totalCost, totalPayAmount, totalRoi)
	avgCost := totalCost/float64(len(reports)) + 0.01
	reportStr := ""
	mainStr := ""
	for _, report := range reports {
		ad := adMap[report.AdID]
		if report.StatCost > avgCost { //大于均值的是主要计划
			mainStr += fmt.Sprintf("主要计划:ID:%d, 消耗:%.2f, ROI:%.2f, 状态:%s\n", report.AdID, report.StatCost, report.PrepayAndPayOrderRoi, ad.Status)
		}
		reportStr += fmt.Sprintf("计划ID:%d, 出价:%.2f, 消耗:%.2f, ROI:%.2f, 成单:%d, 状态:%s\n", report.AdID, ad.DeliverySetting.CPABid, report.StatCost, report.PrepayAndPayOrderRoi, report.PayOrderCount, ad.Status)
	}
	//直播间数据
	liveDatas, err := live_report.GetLiveRoomList(ctx, &live_report.GetLiveRoomListReq{
		AdvertiserID: accountID,
		AwemeID:      2893532936291624,
		DateTime:     time.Now().Add(-8 * time.Hour).Format("2006-01-02"),
		RoomStatus:   live_report.RoomStatusAll,
		AdStatus:     live_report.AdStatusAll,
		Fields:       live_report.RoomMetricsFieldStatCost.All(),
		Page:         1,
		PageSize:     10,
	})
	if err != nil {
		logs.CtxErrorf(ctx, "live_report.GetLiveRoomList failed, err:%v", err)
		lark.SendRoomMessage(ctx, fmt.Sprintf("live_report.GetLiveRoomList failed, err:%v", err))
		return
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
	larkAndLog(ctx, reportStr)
	larkAndLog(ctx, mainStr)
	switch {
	case minutes%10 == 0:
		pCost := totalCost - cost
		if pCost > CostThreshold && cost > 0 {
			//整体降价10%
			larkAndLog(ctx, "整体降价百分之10")
			for _, report := range reports {
				item := adMap[report.AdID]
				err = ad.UpdateAdBid(ctx, accountID, []*ad.Bid{
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
		if pCost < CostMinThreshold && cost > 0 {
			//整体提价10%
			larkAndLog(ctx, "整体提价百分之10")
			for _, report := range reports {
				item := adMap[report.AdID]
				err = ad.UpdateAdBid(ctx, accountID, []*ad.Bid{
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
		cost = totalCost
		for _, report := range reports {
			if report.StatCost < 100.0 && report.PrepayAndPayOrderRoi == 0 { //小于100块而且没有成单,可以先忽略
				continue
			}
			item := adMap[report.AdID]
			if report.PrepayAndPayOrderRoi <= MinROI {
				//半小时后才好做关闭动作
				if minutes <= 30 {
					continue
				}
				//如果已经是危险的了
				_, ok := warningAdMap[item.AdID]
				if ok {
					err = ad.UpdateAdStatus(ctx, &ad.UpdateAdStatusReq{
						AdvertiserID: accountID,
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
				warningAdMap[item.AdID] = item
				larkAndLog(ctx, "设置为危险计划:ID:%d", item.AdID)
			}
			for adID, m := range warningAdMap {
				if m.AdID == item.AdID {
					//如果已经是危险的了
					if report.PrepayAndPayOrderRoi > RecoverRoi {
						//如果已经是危险的了
						delete(warningAdMap, adID)
						larkAndLog(ctx, "重启危险计划:ID:%d,因为ROI恢复了:%.2f", item.AdID, report.PrepayAndPayOrderRoi)
						err = ad.UpdateAdStatus(ctx, &ad.UpdateAdStatusReq{
							AdvertiserID: accountID,
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
}

const MinROI float64 = 2.5
const RecoverRoi float64 = 5

func larkAndLog(ctx context.Context, msg string, args ...interface{}) {
	logs.CtxInfof(ctx, msg, args...)
	lark.SendRoomMessage(ctx, fmt.Sprintf(msg, args...))
}

// 投放配置默认ROI4
const (
	DefaultRoi         float64 = 5
	DefaultBid         float64 = 110
	DefaultFollowUpBid float64 = 0.46
)

const DecreaseBidRatio = 0.92  // 出价降低的比例
const IncreaseRoiRatio = 1.3   // 出价提高的比例
const CostThreshold = 1000.0   // 消耗的阈值, 10分钟1000元
const CostMinThreshold = 200.0 // 消耗的阈值, 10分钟200元
const DefaultBudget = 7000
