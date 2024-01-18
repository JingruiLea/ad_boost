package user_dal

import (
	"context"
	"errors"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
	"gorm.io/gorm"
)

func GetUserByLogin(ctx context.Context, username string, password string) (*model.User, error) {
	db := dal.GetDB(ctx)
	var user model.User
	err := db.Where("username = ? AND passwd = ?", username, password).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.CtxInfof(ctx, "GetUserByLogin not found")
			return nil, nil
		}
		logs.CtxErrorf(ctx, "GetUserByLogin error: %v", err)
		return nil, err
	}
	return &user, nil
}
