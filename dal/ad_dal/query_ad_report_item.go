package ad_dal

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
	"time"
)

type AdReportItemQuery struct {
	AdIDs     []int64
	StartTime time.Time
	EndTime   time.Time
}

func QueryAdReportItem(ctx context.Context, query *AdReportItemQuery) (adReportItem []*model.AdReportItem, err error) {
	db := dal.GetDB(ctx)
	if len(query.AdIDs) > 0 {
		db = db.Where("ad_id in ?", query.AdIDs)
	}
	if !query.StartTime.IsZero() {
		db = db.Where("created_at >= ?", query.StartTime)
	}
	if !query.EndTime.IsZero() {
		db = db.Where("created_at <= ?", query.EndTime)
	}
	err = db.Find(&adReportItem).Error
	if err != nil {
		logs.CtxErrorf(ctx, "QueryAdReportItem db.Find error: %v", err)
		return nil, err
	}
	return adReportItem, nil
}

func GetAdReportByOffset(ctx context.Context, adID int64, offset, limit int) (adReportItem *model.AdReportItem, err error) {
	db := dal.GetDB(ctx)
	adReportItem = &model.AdReportItem{}
	err = db.Model(&model.AdReportItem{}).Where("ad_id = ?", adID).
		Order("id desc").
		Offset(offset).Limit(limit).
		Find(&adReportItem).Error
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdReportByOffset db.Find error: %v", err)
		return nil, err
	}
	return adReportItem, nil
}
