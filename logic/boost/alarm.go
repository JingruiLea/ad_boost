package boost

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/ad_dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/lark"
	"github.com/JingruiLea/ad_boost/logic/ad"
	"github.com/JingruiLea/ad_boost/logic/ad_report"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"time"
)

var AlarmOpen = false

func Alarm(ctx context.Context) {
	AsyncAlarm(ctx, 1748031128935424, time.Minute)
}

func AsyncAlarm(ctx context.Context, accountID int64, interval time.Duration) {
	utils.SafeGo(ctx, func() {
		for range time.Tick(interval) {
			if !AlarmOpen {
				continue
			}
			logs.CtxInfof(ctx, "AsyncAlarm start")
			report := GetReport(ctx, accountID)
			if report != "" {
				lark.SendRoomMessage(ctx, report)
			}
		}
	})
}

func GetReport(ctx context.Context, accountID int64) string {
	return GenReportOnDelivery(ctx, accountID)
}

func GenReportOnDelivery(ctx context.Context, accountID int64) string {
	filter := &ad.Filter{
		MarketingGoal: ttypes.MarketingGoalLivePromGoods,
		Status:        ttypes.AdStatusDeliveryOk,
	}
	return GenReportByFilter(ctx, accountID, filter)
}

func GenReportByFilter(ctx context.Context, accountID int64, filter *ad.Filter) string {
	resp, err := ad.GetAdList(ctx, &ad.GetAdListReq{
		AdvertiserId: accountID,
		Filtering:    filter,
		Page:         1,
		PageSize:     100,
	})
	if err != nil {
		logs.CtxErrorf(ctx, "AsyncAlarm ad.GetAdList error: %v", err)
		return fmt.Sprintf("获取计划列表失败: %v", err)
	}
	if resp == nil || len(resp.List) == 0 {
		logs.CtxInfof(ctx, "AsyncAlarm resp == nil || len(resp.List) == 0")
		return "没有正在投放的计划"
	}
	adMap := make(map[int64]*ad.Ad, len(resp.List))
	adIDs := make([]int64, 0, len(resp.List))
	for _, a := range resp.List {
		adIDs = append(adIDs, a.AdID)
		adMap[a.AdID] = a
	}
	//获取计划报表
	reports, err := ad_report.MGetCommonAdDailyReport(ctx, accountID, adIDs)
	if err != nil {
		logs.CtxErrorf(ctx, "AsyncAlarm ad_report.MGetCommonAdReport error: %v", err)
		return fmt.Sprintf("获取计划报表失败: %v", err)
	}
	if reports == nil || len(reports) == 0 {
		logs.CtxInfof(ctx, "AsyncAlarm reports == nil || len(reports) == 0")
		return "没有获取到计划报表数据"
	}
	models := make([]*model.AdReportItem, 0, len(reports))

	msg := ""
	for _, report := range reports {
		msg += MakeAdReportString(adMap, report)
		models = append(models, report.ToModel(adMap[report.AdID].DeliverySetting.CPABid, adMap[report.AdID].DeliverySetting.ROIGoal))
	}
	err = ad_dal.CreateAdReportItem(ctx, models)
	if err != nil {
		logs.CtxErrorf(ctx, "AsyncAlarm ad_dal.CreateAdReportItem error: %v", err)
	}
	return msg
}

func MakeAdReportString(adMap map[int64]*ad.Ad, report *ad_report.AdReport) string {
	ds := adMap[report.AdID].DeliverySetting
	a := adMap[report.AdID]

	price := ds.CPABid
	priceName := "CPA出价"
	if ds.ExternalAction == ttypes.ExternalActionAdConvertTypeLiveSuccessorderPay &&
		ds.DeepExternalAction == ttypes.DeepExternalActionAdConvertTypeLivePayRoi {
		price = ds.ROIGoal
		priceName = "ROI目标"
	}

	msg := fmt.Sprintf("计划ID: %d, %s:%.2f 名称: %s\n消耗: %.2f, 成单: %d, 支付roi: %.2f, 转化人数: %d, 成交金额: %.2f\n转化率: %.2f, 点击次数: %d, 点击率: %.2f, 展示次数: %d, 平均千次展示费用: %.2f, 新增粉丝数: %d, 转化成本: %.2f\n",
		report.AdID, priceName, price, a.Name, report.StatCost, report.PayOrderCount, report.PrepayAndPayOrderRoi, report.ConvertCnt, report.PayOrderAmount, report.ConvertRate, report.ClickCnt, report.Ctr, report.ShowCnt, report.CpmPlatform, report.DyFollow, report.ConvertCost)

	return msg
}

func GenReportByRedisIDs(ctx context.Context, accountID int64) string {
	ret, err := redis_dal.Get(ctx, fmt.Sprintf("ad_boost:ad_ids:%d", accountID))
	if err != nil {
		logs.CtxErrorf(ctx, "GenReportByRedisIDs redis_dal.Get error: %v", err)
		return fmt.Sprintf("获取计划ID失败: %v", err)
	}
	if ret == "" {
		logs.CtxInfof(ctx, "GenReportByRedisIDs ret == \"\"")
		return "没有获取到计划ID"
	}
	var adIDs []int64
	d := jsoniter.NewDecoder(strings.NewReader(ret))
	d.UseNumber()
	err = d.Decode(&adIDs)
	if err != nil {
		logs.CtxErrorf(ctx, "GenReportByRedisIDs jsoniter.NewDecoder error: %v", err)
		return fmt.Sprintf("获取计划ID失败: %v", err)
	}
	if len(adIDs) == 0 {
		logs.CtxInfof(ctx, "GenReportByRedisIDs len(adIDs) == 0")
		return "没有获取到计划ID"
	}
	filter := &ad.Filter{
		IDs:           adIDs,
		MarketingGoal: ttypes.MarketingGoalLivePromGoods,
	}
	return GenReportByFilter(ctx, accountID, filter)
}
