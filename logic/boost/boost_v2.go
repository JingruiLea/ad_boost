package boost

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/account_dal"
	"github.com/JingruiLea/ad_boost/dal/ad_dal"
	"github.com/JingruiLea/ad_boost/logic/ad"
	"github.com/JingruiLea/ad_boost/logic/ad_report"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"time"
)

func BoostV2MonitorStart(ctx context.Context) {
	allAccounts, err := account_dal.MGetAllAccount(ctx)
	if err != nil {
		logs.CtxErrorf(ctx, "BoostV2MonitorStart account_dal.MGetAllAccount error. err: %s", err.Error())
	}
	for _, account := range allAccounts {
		NewBoostMonitor(ctx, account.AdvertiserID, 5).Start(ctx)
	}
}

type BoostMonitor struct {
	AccountID        int64
	Open             bool
	Mock             bool
	intervalMinCount int64
}

var BoostMonitorMap = make(map[int64]*BoostMonitor)

func (b *BoostMonitor) Start(ctx context.Context) {
	utils.SafeGo(ctx, func() {
		logs.CtxInfof(ctx, "monitor about account %d start", b.AccountID)
		timer := time.Tick(time.Duration(int64(time.Minute) * b.intervalMinCount))
		for range timer {
			b.Loop(ctx)
		}
	})
}

func NewBoostMonitor(ctx context.Context, accountID, intervalMinCount int64) *BoostMonitor {
	ret := &BoostMonitor{
		AccountID:        accountID,
		Open:             true,
		Mock:             false,
		intervalMinCount: intervalMinCount,
	}
	BoostMonitorMap[accountID] = ret
	return ret
}

func (b *BoostMonitor) Stop() {
	b.Open = false
}

func (b *BoostMonitor) Loop(ctx context.Context) {
	if !b.Open {
		logs.CtxInfof(ctx, "%d monitor is closed.", b.AccountID)
		return
	}
	totalAds, err := ad.GetAdListByStatus(ctx, b.AccountID, ttypes.AdStatusDeliveryOk)
	if err != nil {
		logs.CtxErrorf(ctx, "ad.GetAdListByStatus failed, err:%v", err)
		return
	}
	if len(totalAds) == 0 {
		larkAndLog(ctx, "account:%d has no ads in delivery", b.AccountID)
		return
	}
	logs.CtxInfof(ctx, "total ads:%d", len(totalAds))

	adIDs := make([]int64, 0, len(totalAds))
	for _, item := range totalAds {
		adIDs = append(adIDs, item.AdID)
	}
	adMap := make(map[int64]*model.Ad, len(totalAds))
	for _, item := range totalAds {
		adMap[item.AdID] = item
	}
	if len(adIDs) > 100 {
		larkAndLog(ctx, fmt.Sprintf("%d 在投计划数是 %d，超过100个, 请注意", b.AccountID, len(adIDs)))
	}
	reports, err := ad_report.MGetCommonAdDailyReportLarge(ctx, b.AccountID, adIDs)
	if err != nil {
		logs.CtxErrorf(ctx, "ad_report.MGetCommonAdReport failed, err:%v", err)
		return
	}

	adReportItems := make([]*model.AdReportItem, 0, len(reports))
	for _, item := range reports {
		adItem, ok := adMap[item.AdID]
		if !ok {
			logs.CtxErrorf(ctx, "!ok, adID:%d", item.AdID)
			continue
		}
		adReportItems = append(adReportItems, item.ToModel(adItem.DeliverySetting.CPABid, adItem.DeliverySetting.ROIGoal))
	}
	err = ad_dal.CreateAdReportItem(ctx, adReportItems)
	if err != nil {
		logs.CtxErrorf(ctx, "ad_dal.CreateAdReportItem failed, err:%v", err)
		return
	}
	larkAndLog(ctx, fmt.Sprintf("%d 抓取到 %d 个在投计划, 已经存入报表.", b.AccountID, len(adReportItems)))
}
