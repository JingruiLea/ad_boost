package test_dal

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
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
