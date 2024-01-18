package sync

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/ad_dal"
	"github.com/JingruiLea/ad_boost/logic/ad"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
)

func SyncAds(ctx context.Context, advertiserID int64, adID ...int64) error {
	pageSize := 10
	totalPages := 1
	for page := 1; page <= totalPages; page++ {
		req := &ad.GetAdListReq{
			AdvertiserId:     advertiserID,
			RequestAwemeInfo: ttypes.BoolIntTrue,
			Filtering: &ad.Filter{
				MarketingGoal: ttypes.MarketingGoalLivePromGoods,
			},
			Page:     page,
			PageSize: pageSize,
		}
		if len(adID) > 0 {
			req.Filtering.IDs = adID
		}
		resp, err := ad.GetAdList(ctx, req)
		if err != nil {
			logs.CtxErrorf(ctx, "SyncAds ad.GetAdList error: %v", err)
			return err
		}
		if resp == nil || resp.List == nil || len(resp.List) == 0 {
			logs.CtxInfof(ctx, "SyncAds resp == nil || resp.List == nil")
			return nil
		}
		adList := make([]*model.Ad, 0, len(resp.List))
		// Here you can process the ads list, resp.List
		for _, a := range resp.List {
			var reciever model.Ad
			m := reciever.FromBO(a)
			m.AdvertiserID = advertiserID
			adList = append(adList, m)
		}
		err = ad_dal.CreateOrUpdateAds(ctx, adList)
		if err != nil {
			logs.CtxErrorf(ctx, "SyncAds ad_dal.CreateOrUpdateAds error: %v", err)
			return err
		}
		// Update totalPages from the response
		totalPages = resp.PageInfo.TotalPage
	}
	return nil
}
