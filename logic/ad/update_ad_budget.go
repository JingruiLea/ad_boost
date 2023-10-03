package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

const MaxAdsPerUpdate = 10

func MUpdateAdBudgets(ctx context.Context, advertiserID int64, allBudgets []*Budget) error {
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

func UpdateAdBudget(ctx context.Context, advertiserID int64, budgets []*Budget) error {
	var req UpdateAdBudgetReq
	req.AdvertiserId = advertiserID
	req.Data = budgets

	var resp UpdateAdBudgetRespData
	err := httpclient.NewClient().AdPost(ctx, advertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/ad/budget/update/", req, &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "UpdateAdBudget httpclient.NewClient().Post error: %v", err)
		return err
	}
	if len(resp.Errors) > 0 {
		logs.CtxErrorf(ctx, "UpdateAdBudget errors respMap: %s", utils.GetJsonStr(resp))
	}
	return nil
}

type UpdateAdBudgetReq struct {
	AdvertiserId int64     `json:"advertiser_id"`
	Data         []*Budget `json:"data"`
}

type Budget struct {
	AdID   int64   `json:"ad_id"`
	Budget float64 `json:"budget"`
}

type UpdateAdBudgetRespData struct {
	AdIds  []int64       `json:"ad_ids"`
	Errors []interface{} `json:"errors"`
}
