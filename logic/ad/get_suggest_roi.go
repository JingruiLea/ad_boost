package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

type GetSuggestRoiReq struct {
	AdvertiserID   int64                 `json:"advertiser_id"`
	AwemeID        int64                 `json:"aweme_id"`
	MarketingGoal  ttypes.MarketingGoal  `json:"marketing_goal"`
	MarketingScene ttypes.MarketingScene `json:"marketing_scene"`
	ProductID      int64                 `json:"product_id,omitempty"`
	ExternalAction ttypes.ExternalAction `json:"external_action,omitempty"`
	CampaignScene  ttypes.CampaignScene  `json:"campaign_scene,omitempty"`
}

func GetSuggestRoi(ctx context.Context, ad *ttypes.Ad) (float64, error) {
	var req GetSuggestRoiReq
	req.AdvertiserID = ad.AdvertiserID
	req.AwemeID = ad.AwemeID
	req.MarketingGoal = ad.MarketingGoal
	req.MarketingScene = ad.MarketingScene
	req.ExternalAction = ad.DeliverySetting.ExternalAction
	req.CampaignScene = ad.CampaignScene

	mmm := utils.Obj2Map(req)
	var resp GetSuggestRoiRespData
	err := httpclient.NewClient().AdGet(ctx, ad.AdvertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/suggest/roi/goal/", &resp, mmm)
	if err != nil {
		logs.CtxErrorf(ctx, "GetSuggestRoi httpclient.NewClient().Get error: %v", err)
		return 0, err
	}
	fmt.Printf("GetSuggestRoi respMap: %s", utils.GetJsonStr(resp))
	return resp.EcpRoiGoal, err
}

type GetSuggestRoiRespData struct {
	EcpRoiGoal float64 `json:"ecp_roi_goal"`
}
