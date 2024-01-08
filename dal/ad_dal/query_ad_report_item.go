package ad_dal

import (
	"context"
	"errors"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
	"gorm.io/gorm"
	"time"
)

type AdReportItemQuery struct {
	AdIDs     []int64
	StartTime time.Time
	EndTime   time.Time
	Limit     int
	OrderBy   string
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
	if query.Limit > 0 {
		db = db.Limit(query.Limit)
	}
	if query.OrderBy != "" {
		db = db.Order(query.OrderBy)
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

func GetLastAdReport(ctx context.Context, adID int64, lessThan time.Time) (adReportItem model.AdReportItem, err error) {
	db := dal.GetDB(ctx)
	err = db.Model(&model.AdReportItem{}).Where("ad_id = ?", adID).
		Where("created_at < ?", lessThan).
		Last(&adReportItem).Error
	if err != nil {
		logs.CtxErrorf(ctx, "GetLastAdReport db.Find error: %v", err)
		return adReportItem, err
	}
	return adReportItem, nil
}

func MGetLastAdReport(ctx context.Context, adIDs []int64, lessThan time.Time, laterThan time.Time) (adReportItem []*model.AdReportItem, err error) {
	db := dal.GetDB(ctx)
	for _, d := range adIDs {
		var item model.AdReportItem
		err = db.Model(&model.AdReportItem{}).Where("ad_id = ?", d).
			Where("created_at < ?", lessThan).
			Where("created_at > ?", laterThan).
			Last(&item).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			logs.CtxErrorf(ctx, "GetLastAdReport db.Find error: %v", err)
			return nil, err
		}
		adReportItem = append(adReportItem, &item)
	}
	return adReportItem, nil
}
