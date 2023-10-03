package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func MUpdateAdBids(ctx context.Context, advertiserID int64, allBudgets []*Budget) error {
	for i := 0; i < len(allBudgets); i += MaxAdsPerUpdate {
		end := i + MaxAdsPerUpdate
		if end > len(allBudgets) {
			end = len(allBudgets)
		}
		budgets := allBudgets[i:end]

		err := UpdateAdBudget(ctx, advertiserID, budgets)
		if err != nil {
			return fmt.Errorf("UpdateAdBudget failed in batch starting at index %d, err: %v", i, err)
		}
	}
	return nil
}

func UpdateAdBid(ctx context.Context, advertiserID int64, bids []*Bid) error {
	var req UpdateAdBidReq
	req.AdvertiserId = advertiserID
	req.Data = bids

	var resp = make(map[string]interface{})
	err := httpclient.NewClient().AdPost(ctx, advertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/ad/bid/update/", req, &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "UpdateAdBid httpclient.NewClient().Post error: %v", err)
		return err
	}
	fmt.Printf("UpdateAdBid respMap: %s", utils.GetJsonStr(resp))
	return nil
}

type UpdateAdBidReq struct {
	AdvertiserId int64  `json:"advertiser_id"`
	Data         []*Bid `json:"data"`
}

type Bid struct {
	AdId int64   `json:"ad_id"`
	Bid  float64 `json:"bid"`
}
