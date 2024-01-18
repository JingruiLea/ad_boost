package ad_group

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func CreateAdGroup(ctx context.Context, req *CreateAdGroupReq) (*CreateAdGroupRespData, error) {
	var resp CreateAdGroupRespData
	err := httpclient.NewClient().AdPost(ctx, req.AdvertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/campaign/create/", utils.Obj2Map(req), &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "CreateAdGroup httpclient.NewClient().Post error: %v", err)
		return nil, err
	}

	return &resp, nil
}

// CreateAdGroupReq 用于描述广告组的配置信息
type CreateAdGroupReq struct {
	AdvertiserID   int64                 `json:"advertiser_id"`    // 千川广告账户ID，必填
	CampaignName   string                `json:"campaign_name"`    // 广告组名称，必填，1-100个字符，中文字符算2位
	MarketingGoal  ttypes.MarketingGoal  `json:"marketing_goal"`   // 营销目标，必填，允许值：VIDEO_PROM_GOODS 推商品、LIVE_PROM_GOODS 推直播间
	MarketingScene ttypes.MarketingScene `json:"marketing_scene"`  // 广告类型，必填，允许值：FEED 通投广告、SEARCH 搜索广告、SHOPPING_MALL 商城广告
	BudgetMode     ttypes.BudgetMode     `json:"budget_mode"`      // 预算类型，必填，创建后不可修改，允许值：BUDGET_MODE_DAY 日预算、BUDGET_MODE_INFINITE 预算不限
	Budget         float64               `json:"budget,omitempty"` // 广告组预算，必填，当budget_mode为BUDGET_MODE_DAY时，且日预算不少于300元
}

type CreateAdGroupRespData struct {
	CampaignID int64 `json:"campaign_id"`
}
