package account_dal

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
	"gorm.io/gorm/clause"
)

func CreateOrUpdateAweme(ctx context.Context, aweme *model.Aweme) (err error) {
	db := dal.GetDB(ctx)
	err = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "aweme_id"}},
		UpdateAll: true,
	}).Create(&aweme).Error
	if err != nil {
		logs.CtxErrorf(ctx, "CreateOrUpdateAweme db.Create error: %v", err)
		return err
	}
	return nil
}
