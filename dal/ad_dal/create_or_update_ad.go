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
