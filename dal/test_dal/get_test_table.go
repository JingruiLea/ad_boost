package test_dal

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/utils"
	"gorm.io/datatypes"
)

func GetTestByID(ctx context.Context, id int64) (*model.Test, error) {
	db := dal.GetDB(ctx)
	var test model.Test
	err := db.Where("id = ?", id).First(&test).Error
	if err != nil {
		return nil, err
	}
	return &test, nil
}

func CreateAweme(ctx context.Context) error {
	b := []byte(utils.GetJsonStr([]string{"1", "2", "3"}))
	aweme := &model.Aweme{
		ID:                      0,
		AwemeAvatar:             "",
		AwemeHasLivePermission:  false,
		AwemeHasUniProm:         false,
		AwemeHasVideoPermission: false,
		AwemeId:                 0,
		AwemeName:               "",
		AwemeShowId:             "",
		AwemeStatus:             "",
		BindType:                datatypes.JSON(b),
	}
	db := dal.GetDB(ctx)
	err := db.Create(aweme).Error
	if err != nil {
		logs.CtxErrorf(ctx, "CreateAweme db.Create error: %v", err)
		return err
	}
	return nil
}
