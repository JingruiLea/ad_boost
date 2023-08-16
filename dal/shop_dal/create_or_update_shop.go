package shop_dal

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
	"gorm.io/gorm/clause"
)

func CreateOrUpdateShop(ctx context.Context, shop *model.Shop) (err error) {
	db := dal.GetDB(ctx)
	err = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "shop_id"}},
		UpdateAll: true,
	}).Create(&shop).Error
	if err != nil {
		logs.CtxErrorf(ctx, "CreateOrUpdateShop db.Create error: %v", err)
		return err
	}
	return nil
}
