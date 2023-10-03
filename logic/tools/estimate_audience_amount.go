package tools

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func EstimateAudienceAmount(ctx context.Context, adID int64, ad *ttypes.Ad) error {
	req := &EstimateAudienceAmountReq{}
	req = req.FromAd(ad)
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().Get(ctx, "https://api.oceanengine.com/open_api/v1.0/qianchuan/tools/estimate_audience/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "EstimateAudienceAmount httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("EstimateAudienceAmount respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}

type EstimateAudienceAmountReq struct {
	AdvertiserID   int64                 `json:"advertiser_id"`
	MarketingGoal  ttypes.MarketingGoal  `json:"marketing_goal"`
	ExternalAction ttypes.ExternalAction `json:"external_action"`
	AwemeID        int64                 `json:"aweme_id"`
	ProductID      int64                 `json:"product_id"`
	Audience       *ttypes.Audience      `json:"audience"`
}

func (e *EstimateAudienceAmountReq) FromAd(ad *ttypes.Ad) *EstimateAudienceAmountReq {
	e.AdvertiserID = ad.AdvertiserID
	e.MarketingGoal = ad.MarketingGoal
	e.ExternalAction = ad.DeliverySetting.ExternalAction
	e.AwemeID = ad.AwemeID
	e.Audience = ad.Audience
	return e
}
