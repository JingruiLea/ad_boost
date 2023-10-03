package ad_dal

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateOrUpdateAds(ctx context.Context, ad []*model.Ad) (err error) {
	db := dal.GetDB(ctx).Session(&gorm.Session{FullSaveAssociations: true})
	err = db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{
				Name: "ad_id",
			},
		},
		UpdateAll: true,
	}).Create(&ad).Error
	if err != nil {
		logs.CtxErrorf(ctx, "CreateOrUpdateAds db.Create error: %v", err)
		return err
	}
	return nil
}

func CreateOrUpdateDeliverySetting(ctx context.Context, deliverySetting *model.DeliverySetting) (err error) {
	db := dal.GetDB(ctx)
	err = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key_id"}},
		UpdateAll: true,
	}).Create(&deliverySetting).Error
	if err != nil {
		logs.CtxErrorf(ctx, "CreateOrUpdateDeliverySetting db.Create error: %v", err)
		return err
	}
	return nil
}

func CreateOrUpdateAdGroup(ctx context.Context, adGroup []*model.AdGroup) (err error) {
	db := dal.GetDB(ctx)
	err = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ad_group_id"}},
		UpdateAll: true,
	}).Create(&adGroup).Error
	if err != nil {
		logs.CtxErrorf(ctx, "CreateOrUpdateAdGroup db.Create error: %v", err)
		return err
	}
	return nil
}

func CreateOrUpdateAudience(ctx context.Context, audience []*model.Audience) (err error) {
	db := dal.GetDB(ctx)
	err = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "audience_id"}},
		UpdateAll: true,
	}).Create(&audience).Error
	if err != nil {
		logs.CtxErrorf(ctx, "CreateOrUpdateAudience db.Create error: %v", err)
		return err
	}
	return nil
}

func CreateAdReportItem(ctx context.Context, adReportItem []*model.AdReportItem) (err error) {
	db := dal.GetDB(ctx)
	err = db.Create(&adReportItem).Error
	if err != nil {
		logs.CtxErrorf(ctx, "CreateAdReportItem db.Create error: %v", err)
		return err
	}
	return nil
}

func DeleteAdsByAdvertiserID(ctx context.Context, advertiserID int64, adIDs ...int64) (err error) {
	db := dal.GetDB(ctx)
	if len(adIDs) > 0 {
		db = db.Where("ad_id in ?", adIDs)
	}
	err = db.Where("advertiser_id = ?", advertiserID).Delete(&model.Ad{}).Error
	if err != nil {
		logs.CtxErrorf(ctx, "DeleteAdsByAdvertiserID db.Delete error: %v", err)
		return err
	}
	db = db.Session(&gorm.Session{NewDB: true})
	if len(adIDs) > 0 {
		db = db.Where("ad_id in ?", adIDs)
	}
	err = db.Where("advertiser_id = ?", advertiserID).Delete(&model.DeliverySetting{}).Error
	if err != nil {
		logs.CtxErrorf(ctx, "DeleteAdsByAdvertiserID db.Delete error: %v", err)
		return err
	}
	return nil
}

func MGetAdsByAdvertiserID(ctx context.Context, advertiserID int64) (ad []*model.Ad, err error) {
	db := dal.GetDB(ctx)
	err = db.Model(&model.Ad{}).
		Where("advertiser_id = ?", advertiserID).
		Find(&ad).Error
	if err != nil {
		logs.CtxErrorf(ctx, "MGetAdsByAdvertiserID db.Find error: %v", err)
		return nil, err
	}
	adIDs := make([]int64, 0, len(ad))
	adMap := make(map[int64]*model.Ad, len(ad))
	for _, item := range ad {
		adIDs = append(adIDs, item.AdID)
		adMap[item.AdID] = item
	}
	var deliverySetting []*model.DeliverySetting
	err = db.Model(&model.DeliverySetting{}).
		Where("advertiser_id = ?", advertiserID).
		Where("ad_id in ?", adIDs).
		Find(&deliverySetting).Error
	if err != nil {
		logs.CtxErrorf(ctx, "MGetAdsByAdvertiserID db.Find error: %v", err)
		return nil, err
	}
	dsMap := make(map[int64]*model.DeliverySetting, len(deliverySetting))
	for _, item := range deliverySetting {
		dsMap[item.AdID] = item
	}
	for _, item := range ad {
		adMap[item.AdID].DeliverySetting = dsMap[item.AdID]
	}
	return ad, nil
}

func MGetAdByAdIDs(ctx context.Context, adIDs []int64) (adMap map[int64]*model.Ad, err error) {
	var ad []*model.Ad
	adMap = make(map[int64]*model.Ad, len(ad))

	db := dal.GetDB(ctx)
	err = db.Model(&model.Ad{}).
		Where("ad_id in ?", adIDs).
		Find(&ad).Error
	if err != nil {
		logs.CtxErrorf(ctx, "MGetAdByAdIDs db.Find error: %v", err)
		return nil, err
	}
	for _, item := range ad {
		adMap[item.AdID] = item
	}

	var deliverySetting []*model.DeliverySetting
	err = db.Model(&model.DeliverySetting{}).
		Where("ad_id in ?", adIDs).
		Find(&deliverySetting).Error
	if err != nil {
		logs.CtxErrorf(ctx, "MGetAdsByAdvertiserID db.Find error: %v", err)
		return nil, err
	}
	dsMap := make(map[int64]*model.DeliverySetting, len(deliverySetting))
	for _, item := range deliverySetting {
		dsMap[item.AdID] = item
	}
	for _, item := range ad {
		adMap[item.AdID].DeliverySetting = dsMap[item.AdID]
	}

	return adMap, nil
}

func MUpdateAdBudgets(ctx context.Context, advertiserID int64, budget map[int64]float64) error {
	db := dal.GetDB(ctx)
	for k, v := range budget {
		err := db.Model(&model.DeliverySetting{}).Where("ad_id = ?", k).
			Where("advertiser_id = ?", advertiserID).
			Update("budget", v).Error
		if err != nil {
			logs.CtxErrorf(ctx, "MUpdateAdBudgets db.Model error: %v", err)
			return err
		}
	}
	return nil
}
