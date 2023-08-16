package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

type GetSuggestBidReq struct {
	AdvertiserID   int64                 `json:"advertiser_id"`
	AwemeID        int64                 `json:"aweme_id"`
	MarketingGoal  ttypes.MarketingGoal  `json:"marketing_goal"`
	MarketingScene ttypes.MarketingScene `json:"marketing_scene"`
	ProductID      int64                 `json:"product_id,omitempty"`
	ExternalAction ttypes.ExternalAction `json:"external_action,omitempty"`
	CampaignScene  ttypes.CampaignScene  `json:"campaign_scene,omitempty"`
}

func GetSuggestBid(ctx context.Context, ad *ttypes.Ad) (low, high float32, err error) {
	//   "list": [
	//      1748031128935424,
	//    ],
	var req GetSuggestBidReq
	req.AdvertiserID = ad.AdvertiserID
	req.AwemeID = ad.AwemeID
	req.MarketingGoal = ad.MarketingGoal
	req.MarketingScene = ad.MarketingScene
	req.ExternalAction = ad.DeliverySetting.ExternalAction
	req.CampaignScene = ad.CampaignScene

	mmm := utils.Obj2Map(req)
	var resp SuggestBidResp
	err = httpclient.NewClient().Get(ctx, "https://api.oceanengine.com/open_api/v1.0/qianchuan/suggest_bid/", httpclient.CommonHeader, &resp, mmm)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccount httpclient.NewClient().Get error: %v", err)
		return 0, 0, err
	}
	fmt.Printf("GetAdAccount respMap: %s", utils.GetJsonStr(resp))
	//单位为千分之一分, 要转化成元.
	return float32(resp.SuggestBidData.SuggestBidLow) / 100000, float32(resp.SuggestBidData.SuggestBidHigh) / 100000, nil
}

type SuggestBidResp struct {
	ttypes.BaseResp
	SuggestBidData SuggestBidData `json:"data"`
}

type SuggestBidData struct {
	SuggestBidHigh int `json:"suggest_bid_high"`
	SuggestBidLow  int `json:"suggest_bid_low"`
}
