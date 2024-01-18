package admin

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/lark"
	"github.com/JingruiLea/ad_boost/logic/ad"
	"github.com/JingruiLea/ad_boost/model/ttypes"
)

type OnlineAdHandler struct {
}

func (o OnlineAdHandler) GetList(ctx context.Context, filter Filter, range_ Range, sort Sort) (interface{}, int64, error) {
	//accountID, ok := filter["account_id"].(int64)
	//if !ok {
	//	return nil, 0, nil
	//}
	accountID := 1784698853978186
	totalAds, err := ad.GetAdListByStatus(ctx, int64(accountID), ttypes.AdStatusDeliveryOk, 0, nil)
	if err != nil {
		logs.CtxErrorf(ctx, "ad.GetAdListByStatus failed, err:%v", err)
		lark.SendRoomMessage(ctx, fmt.Sprintf("ad.GetAdListByStatus failed, err:%v", err))
		return nil, 0, err
	}
	for _, aditem := range totalAds {
		aditem.ID = uint(aditem.AdID)
	}
	return totalAds, int64(len(totalAds)), nil
}

func (o OnlineAdHandler) GetOne(ctx context.Context, id int64) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (o OnlineAdHandler) GetMany(ctx context.Context, filter GetManyFilter) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (o OnlineAdHandler) GetManyReference(ctx context.Context, filter GetManyFilter) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (o OnlineAdHandler) Create(ctx context.Context, data map[string]interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (o OnlineAdHandler) Update(ctx context.Context, id int64, data interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (o OnlineAdHandler) Delete(ctx context.Context, id int64) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}
