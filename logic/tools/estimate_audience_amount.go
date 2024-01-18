package tools

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/audience_dal"
	"github.com/JingruiLea/ad_boost/model/bo"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func EstimateAudienceAmountFromAd(ctx context.Context, ad *bo.CreateAd) (*EstimateAudienceAmountData, error) {
	req := &EstimateAudienceAmountReq{}
	req = req.FromAd(ad)
	return EstimateAudienceAmount(ctx, req)
}

// 把http方法抽象出来
func EstimateAudienceAmount(ctx context.Context, req *EstimateAudienceAmountReq) (*EstimateAudienceAmountData, error) {
	var resp EstimateAudienceAmountData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://api.oceanengine.com/open_api/v1.0/qianchuan/tools/estimate_audience/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "EstimateAudienceAmount httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	fmt.Printf("EstimateAudienceAmount respMap: %s", utils.GetJsonStr(resp))
	return &resp, nil
}

func EstimateAudienceAmountByAudienceID(ctx context.Context, audienceID int64, awemeID int64) (*EstimateAudienceAmountData, error) {
	audience, err := audience_dal.GetAudienceByAudienceID(ctx, audienceID)
	if err != nil {
		logs.CtxErrorf(ctx, "EstimateAudienceAmountByAudienceID audience_dal.GetAudienceByAudienceID error: %v", err)
		return nil, err
	}
	req := &EstimateAudienceAmountReq{
		AdvertiserID:   audience.AccountID,
		MarketingGoal:  ttypes.MarketingGoalLivePromGoods,
		ExternalAction: ttypes.ExternalActionAdConvertTypeLiveSuccessorderPay,
		AwemeID:        awemeID,
		Audience: &bo.Audience{
			AudienceMode:  ttypes.AudienceModeCustom,
			OrientationID: audienceID,
		},
	}
	return EstimateAudienceAmount(ctx, req)
}

type EstimateAudienceAmountReq struct {
	AdvertiserID   int64                 `json:"advertiser_id"`
	MarketingGoal  ttypes.MarketingGoal  `json:"marketing_goal"`
	ExternalAction ttypes.ExternalAction `json:"external_action"`
	AwemeID        int64                 `json:"aweme_id,omitempty"`
	ProductID      int64                 `json:"product_id,omitempty"`
	Audience       *bo.Audience          `json:"audience"`
}

func (e *EstimateAudienceAmountReq) FromAd(ad *bo.CreateAd) *EstimateAudienceAmountReq {
	e.AdvertiserID = ad.AdvertiserID
	e.MarketingGoal = ad.MarketingGoal
	e.ExternalAction = ad.DeliverySetting.ExternalAction
	e.AwemeID = ad.AwemeID
	e.Audience = ad.Audience
	return e
}

type EstimateAudienceAmountData struct {
	CrowdCoverTotal int64 `json:"crowd_cover_total"`
	ShowCntTotal    int64 `json:"show_cnt_total"`
}
