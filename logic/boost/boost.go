package boost

import (
	"context"
	"errors"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/ad_dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/lark"
	"github.com/JingruiLea/ad_boost/logic/ad"
	"github.com/JingruiLea/ad_boost/logic/ad_report"
	"github.com/JingruiLea/ad_boost/logic/boost/sync"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Roi float64

var boostOn bool

var fakeRun bool

var boostStart *time.Time

var cost float64

// warning_ad_map
var warningAdMap = make(map[int64]*model.Ad)

func Init() {
	Alarm(context.Background())
	lark.RegisterTextHandler(context.Background(), func(msg string) string {
		logs.CtxInfof(context.Background(), "收到消息:%s", msg)
		switch {
		case strings.HasPrefix(msg, "boost init"):
			err := BoostInit(context.Background(), 1748031128935424)
			if err != nil {
				return fmt.Sprintf("boost init failed, err:%v", err)
			}
			return "boost init done"
		case strings.HasPrefix(msg, "boost start"):
			//parse start time
			startTimeStr := strings.TrimPrefix(msg, "boost start")
			startTimeStr = strings.TrimSpace(startTimeStr)
			if startTimeStr == "" {
				return "boost start failed, start time is empty"
			}
			BoostStart(context.Background(), 1748031128935424, startTimeStr)
			return "boost start"
		case strings.HasPrefix(msg, "boost stop"):
			boostOn = false
			return "boost stop"
		case strings.HasPrefix(msg, "boost status"):
			return fmt.Sprintf("boost status: %v", boostOn)
		case strings.HasPrefix(msg, "boost fake run"):
			fakeRun = true
			return "boost fake run"
		case strings.HasPrefix(msg, "boost fake stop"):
			fakeRun = false
			//实时报表
		case strings.HasPrefix(msg, "boost report start"):
			AlarmOpen = true
			return "boost report"
		case strings.HasPrefix(msg, "boost report stop"):
			AlarmOpen = false
		case strings.HasPrefix(msg, "更新出价"):
			//更新出价 ID 出价
			//更新出价 123 0.5
			adArr := strings.Split(msg, " ")
			if len(adArr) != 3 {
				return "更新出价失败,参数不对"
			}
			adID := utils.Str2I64(adArr[1])
			if adID == 0 {
				return "更新出价失败,参数不对"
			}
			bid := utils.Str2Float64(adArr[2])
			if bid == 0 {
				return "更新出价失败,参数不对"
			}
			err := ad.UpdateAdBid(context.Background(), 1748031128935424, []*ad.Bid{
				{
					AdId: adID,
					Bid:  bid,
				},
			})
			if err != nil {
				return fmt.Sprintf("更新出价失败, err:%v", err)
			}

		case strings.HasPrefix(msg, "停止"):
			//停止 ID
			adArr := strings.Split(msg, " ")
			if len(adArr) != 2 {
				return "停止失败,参数不对"
			}
			adID := utils.Str2I64(adArr[1])
			if adID == 0 {
				return "停止失败,参数不对"
			}
			err := ad.UpdateAdStatus(context.Background(), &ad.UpdateAdStatusReq{
				AdvertiserID: 1748031128935424,
				AdIDs:        []int64{adID},
				OptStatus:    ad.OptStatusDisable,
			})
			if err != nil {
				return fmt.Sprintf("停止失败, err:%v", err)
			}

		case strings.HasPrefix(msg, "启动"):
			//启动 ID
			adArr := strings.Split(msg, " ")
			if len(adArr) != 2 {
				return "启动失败,参数不对"
			}
			adID := utils.Str2I64(adArr[1])
			if adID == 0 {
				return "启动失败,参数不对"
			}
			err := ad.UpdateAdStatus(context.Background(), &ad.UpdateAdStatusReq{
				AdvertiserID: 1748031128935424,
				AdIDs:        []int64{adID},
				OptStatus:    ad.OptStatusEnable,
			})
			if err != nil {
				return fmt.Sprintf("启动失败, err:%v", err)
			}
		}
		return ""
	})
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
	//把所有广告预算改到300, 防止爆炸
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
	err = ad.MUpdateAdStatus(ctx, accountID, adIDs, ad.OptStatusEnable)
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

func BoostStart(ctx context.Context, accountID int64, startTimeStr string) {
	larkAndLog(ctx, "开始启动定时监测,V2")
	BoostLoop(ctx, accountID, startTimeStr)
}

var startTime time.Time

func BoostLoop(ctx context.Context, accountID int64, startTimeStr string) {
	boostOn = true
	var tickCount int
	utils.SafeGo(ctx, func() {
		for range time.Tick(time.Minute * 1) {
			if !boostOn {
				larkAndLog(ctx, "定时监测关闭")
				tickCount = 0
				return
			}
			tickCount++
			_, err := time.Parse("2006-01-02 15:04:05", startTimeStr)
			if err != nil {
				logs.CtxErrorf(ctx, "time.Parse error: %v", err)
				lark.SendRoomMessage(ctx, fmt.Sprintf("time.Parse error: %v", err))
				return
			}
			//正在投放
			resp, err := ad.GetAdListByStatus(ctx, accountID, ttypes.AdStatusDeliveryOk)
			if err != nil {
				logs.CtxErrorf(ctx, "ad.GetAdListByStatus failed, err:%v", err)
				lark.SendRoomMessage(ctx, fmt.Sprintf("ad.GetAdListByStatus failed, err:%v", err))
				continue
			}

			if len(resp) == 0 {
				larkAndLog(ctx, "没有正在投放的计划")
				boostStart = nil
				continue
			}
			var totalAds []*model.Ad
			//暂停的计划
			respPause, err := ad.GetAdListByStatus(ctx, accountID, ttypes.AdStatusDisable)
			if err != nil {
				logs.CtxErrorf(ctx, "ad.GetAdListByStatus failed, err:%v", err)
				lark.SendRoomMessage(ctx, fmt.Sprintf("ad.GetAdListByStatus failed, err:%v", err))
				continue
			}
			if len(respPause) > 0 {
				totalAds = append(totalAds, respPause...)
			}
			var pausedIDs []int64
			larkAndLog(ctx, "找到%d个正在投放的计划", len(resp))
			if boostStart == nil {
				now := time.Now()
				boostStart = &now
				for _, m := range respPause {
					pausedIDs = append(pausedIDs, m.AdID)
				}
				//启动所有暂停的计划
				err = ad.UpdateAdStatus(ctx, &ad.UpdateAdStatusReq{
					AdvertiserID: accountID,
					AdIDs:        pausedIDs,
					OptStatus:    ad.OptStatusEnable,
				})
				if err != nil {
					logs.CtxErrorf(ctx, "ad.UpdateAdStatus failed, err:%v", err)
					lark.SendRoomMessage(ctx, fmt.Sprintf("ad.UpdateAdStatus failed, err:%v", err))
					continue
				}
			}
			adIDs := make([]int64, 0, len(resp))
			for _, item := range resp {
				adIDs = append(adIDs, item.AdID)
			}

			adMap, err := ad_dal.MGetAdByAdIDs(ctx, adIDs)
			if err != nil {
				logs.CtxErrorf(ctx, "ad_dal.MGetAdByAdIDs failed, err:%v", err)
				lark.SendRoomMessage(ctx, fmt.Sprintf("ad_dal.MGetAdByAdIDs failed, err:%v", err))
				return
			}

			reports, err := ad_report.MGetCommonAdReport(ctx, accountID, adIDs)
			if err != nil {
				logs.CtxErrorf(ctx, "ad_report.MGetCommonAdReport failed, err:%v", err)
				lark.SendRoomMessage(ctx, fmt.Sprintf("ad_report.MGetCommonAdReport failed, err:%v", err))
				continue
			}
			adReportItems := make([]*model.AdReportItem, 0, len(reports))
			for _, item := range reports {
				adReportItems = append(adReportItems, item.ToModel(adMap[item.AdID].DeliverySetting.CpaBid, adMap[item.AdID].DeliverySetting.RoiGoal))
			}
			err = ad_dal.CreateAdReportItem(ctx, adReportItems)
			if err != nil {
				logs.CtxErrorf(ctx, "ad_dal.CreateAdReportItem failed, err:%v", err)
				lark.SendRoomMessage(ctx, fmt.Sprintf("ad_dal.CreateAdReportItem failed, err:%v", err))
				continue
			}
			minutes := int(time.Since(*boostStart).Minutes())
			if minutes < 0 {
				logs.CtxErrorf(ctx, "time.Since(*boostStart).Minutes() < 0")
				lark.SendRoomMessage(ctx, fmt.Sprintf("time.Since(*boostStart).Minutes() < 0"))
				continue
			}
			larkAndLog(ctx, "开始第%d分钟的监测", minutes)
			totalPayAmount := float64(0)
			totalCost := float64(0)
			totalRoi := float64(0)
			for _, report := range reports {
				totalPayAmount += report.PayOrderAmount
				totalCost += report.StatCost
			}
			totalRoi = utils.RoundFloat(totalPayAmount/totalCost, 2)
			larkAndLog(ctx, "总消耗:%.2f, 总成交金额:%.2f, 总ROI:%.2f", totalCost, totalPayAmount, totalRoi)
			avgCost := totalCost/float64(len(reports)) + 0.01
			reportStr := ""
			for _, report := range reports {
				if report.StatCost > avgCost { //大于均值的是主要计划
					larkAndLog(ctx, "主要计划:ID:%d, 消耗:%.2f, ROI:%.2f", report.AdID, report.StatCost, report.PrepayAndPayOrderRoi)
				}
				reportStr += fmt.Sprintf("计划ID:%d, 消耗:%.2f, ROI:%.2f, 成单:%d", report.AdID, report.StatCost, report.PrepayAndPayOrderRoi, report.PayOrderCount)
			}

			switch {
			case minutes%10 == 0:
				pCost := totalCost - cost
				if pCost > CostThreshold {
					//整体降价10%
					larkAndLog(ctx, "整体降价10%")
					for _, report := range reports {
						item := adMap[report.AdID]
						err = ad.UpdateAdBudget(ctx, accountID, []*ad.Budget{
							{
								AdID:   report.AdID,
								Budget: utils.RoundFloat(item.DeliverySetting.CpaBid*0.9, 2),
							},
						})
						if err != nil {
							logs.CtxErrorf(ctx, "ad.UpdateAdBudget failed, err:%v", err)
							lark.SendRoomMessage(ctx, fmt.Sprintf("ad.UpdateAdBudget failed, err:%v", err))
							continue
						}
					}
				}
				if pCost < CostMinThreshold {
					//整体提价10%
					larkAndLog(ctx, "整体提价10%")
					for _, report := range reports {
						item := adMap[report.AdID]
						err = ad.UpdateAdBudget(ctx, accountID, []*ad.Budget{
							{
								AdID:   report.AdID,
								Budget: utils.RoundFloat(item.DeliverySetting.CpaBid*1.1, 2),
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
								OptStatus:    ad.OptStatusDisable,
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
									OptStatus:    ad.OptStatusEnable,
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
	})
}

const MinROI float64 = 2.5
const RecoverRoi float64 = 5

func GetReportBefore3MinAnd6Min(ctx context.Context, adID int64) (min3 *model.AdReportItem, min6 *model.AdReportItem, err error) {
	min3, err = ad_dal.GetAdReportByOffset(ctx, adID, 3, 1)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.CtxInfof(ctx, "还么有3分钟的数据")
			return nil, nil, nil
		}
		logs.CtxErrorf(ctx, "ad_dal.GetAdReportByOffset failed, err:%v", err)
		return nil, nil, err
	}
	min6, err = ad_dal.GetAdReportByOffset(ctx, adID, 6, 1)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.CtxInfof(ctx, "还么有6分钟的数据")
			return min3, nil, nil
		}
		logs.CtxErrorf(ctx, "ad_dal.GetAdReportByOffset failed, err:%v", err)
		return nil, nil, err
	}
	return min3, min6, nil
}

func decreaseBidAndSync(ctx context.Context, accountID int64, item *model.Ad) error {
	switch {
	case item.IsRoi():
		// 如果消耗超过阈值，并且还没有提高过ROI，提升广告计划的ROI
		newRoi := utils.RoundFloat(item.DeliverySetting.RoiGoal*IncreaseRoiRatio, 2)
		item.DeliverySetting.RoiGoal = newRoi
		if !fakeRun {
			err := ad.UpdateAdRoiGoal(ctx, accountID, []*ad.RoiGoal{
				{
					AdId:    item.AdID,
					RoiGoal: newRoi,
				},
			})
			if err != nil {
				logs.CtxErrorf(ctx, "ad.UpdateAdRoiGoal failed, err:%v", err)
				return err
			}
		}
	case item.IsCpa():
		// 如果消耗超过阈值，并且还没有降低过出价，降低广告计划的出价
		newBid := utils.RoundFloat(item.DeliverySetting.CpaBid*DecreaseBidRatio, 2)
		item.DeliverySetting.CpaBid = newBid
		if !fakeRun {
			err := ad.UpdateAdBid(ctx, accountID, []*ad.Bid{
				{
					AdId: item.AdID,
					Bid:  newBid,
				},
			})
			if err != nil {
				logs.CtxErrorf(ctx, "ad.UpdateAdBid failed, err:%v", err)
				return err
			}
		}
	}
	// 记录该广告计划的出价已经被降低
	err := markBidIncreased(ctx, item.AdID)
	if err != nil {
		logs.CtxErrorf(ctx, "markBidIncreased failed, err:%v", err)
		return err
	}
	err = sync.SyncAds(ctx, accountID, item.AdID)
	if err != nil {
		logs.CtxErrorf(ctx, "sync.SyncAds failed, err:%v", err)
		return err
	}
	return nil
}

func markBidIncreased(ctx context.Context, id int64) error {
	return redis_dal.GetRedisClient().Set(ctx, fmt.Sprintf("ad_boost:bid_increased:%d", id), 1, time.Hour*5).Err()
}

func hasDncreasedBid(ctx context.Context, id int64) (bool, error) {
	key := fmt.Sprintf("ad_boost:bid_increased:%d", id)
	res, err := redis_dal.GetRedisClient().Exists(ctx, key).Result()
	if err != nil {
		logs.CtxErrorf(ctx, "error checking existence of key %s: %v", key, err)
		return false, fmt.Errorf("error checking existence of key %s: %v", key, err)
	}
	return res > 0, nil
}

func larkAndLog(ctx context.Context, msg string, args ...interface{}) {
	logs.CtxInfof(ctx, msg, args...)
	lark.SendRoomMessage(ctx, fmt.Sprintf(msg, args...))
}

func buildAudience(ctx context.Context, adID int64) error {
	//获取人群包
	//创建人群包
	//创建人群包定向
	return nil
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
const CostMinThreshold = 100.0 // 消耗的阈值, 10分钟100元
const DefaultBudget = 1000
