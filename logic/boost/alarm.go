package boost

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/lark"
	"github.com/JingruiLea/ad_boost/logic/ad"
	"github.com/JingruiLea/ad_boost/logic/ad_report"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"strings"
	"time"
)

var AlarmOpen = false

func Init() {
	lark.RegisterTextHandler(context.Background(), func(msg string) string {
		if strings.Contains(msg, "关闭实时报表") {
			AlarmOpen = false
			return "实时报表功能已关闭"
		}
		if strings.Contains(msg, "实时报表") {
			AlarmOpen = true
			return "实时报表功能已开启, 每分钟播报."
		}
		if strings.Contains(msg, "报表") {
			report := GenReport(context.Background(), 1748031128935424)
			if strings.Contains(msg, "没有") {
				AlarmOpen = false
				return report
			}
			return report
		}
		return ""
	})
	Alarm(context.Background())
}

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
			report := GenReport(ctx, accountID)
			if report != "" {
				lark.SendRoomMessage(ctx, report)
			}
		}
	})
}

func GenReport(ctx context.Context, accountID int64) string {
	resp, err := ad.GetAdList(ctx, &ad.GetAdListReq{
		AdvertiserId: accountID,
		Filtering: &ad.Filter{
			MarketingGoal: ttypes.MarketingGoalLivePromGoods,
			Status:        ad.AdStatusDeliveryOk, //获取投放中计划
		},
		Page:     1,
		PageSize: 100,
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
	reports, err := ad_report.MGetCommonAdReport(ctx, accountID, adIDs)
	if err != nil {
		logs.CtxErrorf(ctx, "AsyncAlarm ad_report.MGetCommonAdReport error: %v", err)
		return fmt.Sprintf("获取计划报表失败: %v", err)
	}
	if reports == nil || len(reports) == 0 {
		logs.CtxInfof(ctx, "AsyncAlarm reports == nil || len(reports) == 0")
		return "没有获取到计划报表数据"
	}
	msg := ""
	for _, report := range reports {
		msg += MakeAdReportString(adMap, report)
	}
	return msg
}

func MakeAdReportString(adMap map[int64]*ad.Ad, report *ad_report.AdReport) string {
	ds := adMap[report.AdID].DeliverySetting
	ad := adMap[report.AdID]

	price := ds.CpaBid
	priceName := "CPA出价"
	if ds.ExternalAction == ttypes.ExternalActionAdConvertTypeLiveSuccessorderPay &&
		ds.DeepExternalAction == ttypes.DeepExternalActionAdConvertTypeLivePayRoi {
		price = ds.RoiGoal
		priceName = "ROI目标"
	}

	msg := fmt.Sprintf("计划ID: %d, %s:%.2f 名称: %s\n消耗: %.2f, 成单: %d, 支付roi: %.2f, 转化人数: %d, 成交金额: %.2f\n转化率: %.2f, 点击次数: %d, 点击率: %.2f, 展示次数: %d, 平均千次展示费用: %.2f, 新增粉丝数: %d, 转化成本: %.2f\n",
		report.AdID, priceName, price, ad.Name, report.StatCost, report.PayOrderCount, report.PrepayAndPayOrderRoi, report.ConvertCnt, report.PayOrderAmount, report.ConvertRate, report.ClickCnt, report.Ctr, report.ShowCnt, report.CpmPlatform, report.DyFollow, report.ConvertCost)

	return msg
}
