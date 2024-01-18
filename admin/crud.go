package admin

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
)

type CrudHandler[T any] struct {
}

func (a CrudHandler[T]) GetList(ctx context.Context, filter Filter, range_ Range, sort Sort) (interface{}, int64, error) {
	db := dal.GetDB(ctx)
	var result []*T
	var tIns T
	var start = range_.Start
	var end = range_.End
	err := db.Model(&tIns).Offset(start).Limit(end - start + 1).Order(sort.Field + " " + string(sort.Direction)).Find(&result).Error
	if err != nil {
		logs.CtxErrorf(ctx, "GetList error: %v", err)
		return nil, 0, err
	}
	var total int64
	err = db.Model(&tIns).Count(&total).Error
	if err != nil {
		logs.CtxErrorf(ctx, "GetList error: %v", err)
		return nil, 0, err
	}
	return result, total, err
}

func (a CrudHandler[T]) GetOne(ctx context.Context, id int64) (interface{}, error) {
	db := dal.GetDB(ctx)
	var result T
	err := db.Model(&result).Where("id = ?", id).First(&result).Error
	if err != nil {
		logs.CtxErrorf(ctx, "GetOne error: %v", err)
		return nil, err
	}
	return &result, err
}

func (a CrudHandler[T]) GetMany(ctx context.Context, filter GetManyFilter) (interface{}, error) {
	db := dal.GetDB(ctx)
	var result []T
	err := db.Model(&result).Where("id in (?)", filter.Ids).Find(&result).Error
	if err != nil {
		logs.CtxErrorf(ctx, "GetMany error: %v", err)
		return nil, err
	}
	return result, err
}

func (a CrudHandler[T]) GetManyReference(ctx context.Context, filter GetManyFilter) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (a CrudHandler[T]) Create(ctx context.Context, data map[string]interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (a CrudHandler[T]) Update(ctx context.Context, id int64, data interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (a CrudHandler[T]) Delete(ctx context.Context, id int64) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}
