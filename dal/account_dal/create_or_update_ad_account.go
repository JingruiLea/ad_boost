package account_dal

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
	"gorm.io/gorm/clause"
)

func CreateOrUpdateAdAccount(ctx context.Context, adAccount *model.Advertiser) (err error) {
	db := dal.GetDB(ctx)
	err = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "advertiser_id"}},
		UpdateAll: true,
	}).Create(&adAccount).Error
	if err != nil {
		logs.CtxErrorf(ctx, "CreateOrUpdateAdAccount db.Create error: %v", err)
		return err
	}
	return nil
}

func GetAdAccountByAccountID(ctx context.Context, accountID int64) (*model.Advertiser, error) {
	db := dal.GetDB(ctx)
	var adAccount model.Advertiser
	err := db.Where("advertiser_id = ?", accountID).First(&adAccount).Error
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccountByAccountID db.Where error: %v", err)
		return nil, err
	}
	return &adAccount, nil
}
