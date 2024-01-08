package account_dal

import (
	"context"
	"errors"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
	"gorm.io/gorm"
)

func CreateOrUpdateAweme(ctx context.Context, aweme *model.Aweme) (err error) {
	db := dal.GetDB(ctx)
	var ret int64
	err = db.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&aweme).Where("aweme_id = ? AND account_id = ?", aweme.AwemeId, aweme.AccountID).Count(&ret).Error
		if err != nil {
			logs.CtxErrorf(ctx, "CreateOrUpdateAweme db.Where error: %v", err)
			return err
		}
		if ret > 1 {
			logs.CtxErrorf(ctx, "CreateOrUpdateAweme db.Where ret > 1")
			return errors.New("CreateOrUpdateAweme db.Where ret > 1")
		}
		//如果存在则更新
		if ret > 0 {
			err = tx.Model(&model.Aweme{}).Where("aweme_id = ? AND account_id = ?", aweme.AwemeId, aweme.AccountID).Updates(aweme).Error
			if err != nil {
				logs.CtxErrorf(ctx, "CreateOrUpdateAweme tx.Model error: %v", err)
				return err
			}
		} else {
			err = tx.Create(aweme).Error
			if err != nil {
				logs.CtxErrorf(ctx, "CreateOrUpdateAweme tx.Create error: %v", err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		logs.CtxErrorf(ctx, "CreateOrUpdateAweme db.Transaction error: %v", err)
		return err
	}
	return nil
}
