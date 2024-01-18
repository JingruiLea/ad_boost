package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
	"log"
)

func MUpdateAdStatus(ctx context.Context, advertiserID int64, allAdIDs []int64, optStatus ttypes.OptStatus) error {
	var i int
	for i = 0; i < len(allAdIDs); i += MaxAdsPerUpdate {
		end := i + MaxAdsPerUpdate
		if end > len(allAdIDs) {
			end = len(allAdIDs)
		}

		// Create request object
		req := &UpdateAdStatusReq{
			AdvertiserID: advertiserID,
			AdIDs:        allAdIDs[i:end],
			OptStatus:    optStatus,
		}

		// Call the UpdateAdStatus method with the request
		err := UpdateAdStatus(ctx, req)
		if err != nil {
			// If an error occurs, log it and return it
			log.Printf("Error updating ad status: %v", err)
			return err
		}
	}
	return nil
}

func UpdateAdStatus(ctx context.Context, req *UpdateAdStatusReq) error {
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().AdPost(ctx, req.AdvertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/ad/status/update/", req, &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "UpdateAdStatus httpclient.NewClient().Post error: %v", err)
		return err
	}
	fmt.Printf("UpdateAdStatus respMap: %s", utils.GetJsonStr(resp))
	return nil
}

type UpdateAdStatusReq struct {
	AdvertiserID int64            `json:"advertiser_id"`
	AdIDs        []int64          `json:"ad_ids"` //最多10个
	OptStatus    ttypes.OptStatus `json:"opt_status"`
}
