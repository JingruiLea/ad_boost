package sync

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/audience_dal"
	"github.com/JingruiLea/ad_boost/logic/ad_group"
	"github.com/JingruiLea/ad_boost/logic/audience"
	"github.com/JingruiLea/ad_boost/model"
)

func SyncAudiencePackage(ctx context.Context, advertiserID int64, filter *ad_group.Filter) error {
	var req audience.GetAudiencePackageListReq
	req.Page = 1
	req.PageSize = 10
	req.AdvertiserID = advertiserID
	for {
		resp, err := audience.GetAudiencePackageList(ctx, &req)
		if err != nil {
			logs.CtxErrorf(ctx, "SyncAdGroup ad_group.GetAdGroupList error: %v", err)
			return err
		}
		if resp == nil || resp.List == nil {
			logs.CtxInfof(ctx, "SyncAdGroup resp == nil || resp.List == nil")
			return nil
		}
		audiences := make([]*model.Audience, 0, len(resp.List))
		for _, audiencePacakge := range resp.List {
			var reciever model.Audience
			audiences = append(audiences, reciever.FromBO(audiencePacakge, advertiserID))
		}
		err = audience_dal.CreateOrUpdateAudience(ctx, audiences)
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
