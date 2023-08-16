package sync

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/ad_dal"
	"github.com/JingruiLea/ad_boost/logic/ad"
	"github.com/JingruiLea/ad_boost/model"
)

func SyncAds(ctx context.Context, advertiserID int64) error {
	pageSize := 10
	totalPages := 1

	for page := 1; page <= totalPages; page++ {
		resp, err := ad.GetAdList(ctx, &ad.GetAdListReq{
			AdvertiserId:     advertiserID,
			RequestAwemeInfo: ad.AwemeInfoNoInclude,
			Filtering: ad.Filter{
				MarketingGoal: ad.MarketingGoalFilterLivePromGoods,
			},
			Page:     page,
			PageSize: pageSize,
		})
		if err != nil {
			logs.CtxErrorf(ctx, "SyncAds ad.GetAdList error: %v", err)
			return err
		}
		if resp == nil || resp.List == nil {
			logs.CtxInfof(ctx, "SyncAds resp == nil || resp.List == nil")
			return nil
		}
		adList := make([]*model.Ad, 0, len(resp.List))
		// Here you can process the ads list, resp.List
		for _, a := range resp.List {
			adList = append(adList, a.ToModel())
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
