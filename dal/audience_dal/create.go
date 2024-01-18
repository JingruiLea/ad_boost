package audience_dal

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
	"gorm.io/gorm/clause"
)

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

func GetAudienceByAudienceID(ctx context.Context, audienceID int64) (audience *model.Audience, err error) {
	db := dal.GetDB(ctx)
	err = db.Where("audience_id = ?", audienceID).First(&audience).Error
	if err != nil {
		logs.CtxErrorf(ctx, "GetAudienceByAudienceID db.Where error: %v", err)
		return nil, err
	}
	return audience, nil
}
