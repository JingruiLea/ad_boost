package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func UpdateAdBudget(ctx context.Context, advertiserID int64, budgets []*Budget) error {
	var req UpdateAdBudgetReq
	req.AdvertiserId = advertiserID
	req.Data = budgets

	var resp UpdateAdBudgetResp
	err := httpclient.NewClient().Post(ctx, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/ad/budget/update/", httpclient.CommonHeader, req, &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "UpdateAdBudget httpclient.NewClient().Post error: %v", err)
		return err
	}
	fmt.Printf("UpdateAdBudget respMap: %s", utils.GetJsonStr(resp))
	if resp.Code != 0 {
		logs.CtxInfof(ctx, "UpdateAdBudget resp.Code != 0")
		return fmt.Errorf("UpdateAdBudget resp.Code != 0")
	}
	return nil
}

type UpdateAdBudgetReq struct {
	AdvertiserId int64     `json:"advertiser_id"`
	Data         []*Budget `json:"data"`
}

type Budget struct {
	AdId   int64 `json:"ad_id"`
	Budget int64 `json:"budget"`
}

type UpdateAdBudgetResp struct {
	ttypes.BaseResp
	Data *UpdateAdBudgetRespData `json:"data"`
}

type UpdateAdBudgetRespData struct {
	AdIds  []int64       `json:"ad_ids"`
	Errors []interface{} `json:"errors"`
}
