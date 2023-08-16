package sync

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/ad_dal"
	"github.com/JingruiLea/ad_boost/logic/ad_group"
	"github.com/JingruiLea/ad_boost/model"
)

func SyncAdGroup(ctx context.Context, advertiserID int64, filter *ad_group.Filter) error {
	var req ad_group.GetAdGroupListReq
	req.Page = 1
	req.PageSize = 10
	req.AdvertiserID = advertiserID
	req.Filter = filter
	for {
		resp, err := ad_group.GetAdGroupList(ctx, &req)
		if err != nil {
			logs.CtxErrorf(ctx, "SyncAdGroup ad_group.GetAdGroupList error: %v", err)
			return err
		}
		if resp == nil || resp.List == nil {
			logs.CtxInfof(ctx, "SyncAdGroup resp == nil || resp.List == nil")
			return nil
		}
		adGroupList := make([]*model.AdGroup, 0, len(resp.List))
		for _, campaign := range resp.List {
			adGroupList = append(adGroupList, campaign.ToModel())
		}
		err = ad_dal.CreateOrUpdateAdGroup(ctx, adGroupList)
		if err != nil {
			logs.CtxErrorf(ctx, "SyncAdGroup ad_dal.CreateOrUpdateAdGroup error: %v", err)
			return err
		}
		if resp.PageInfo.TotalPage <= req.Page {
			break
		}
		req.Page++
	}
	return nil
}
