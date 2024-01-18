package ad_group

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func UpdateAdGroup(ctx context.Context, req *UpdateAdGroupReq) (*UpdateAdGroupRespData, error) {
	var resp UpdateAdGroupRespData
	err := httpclient.NewClient().AdPost(ctx, req.AdvertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/campaign/update/", utils.Obj2Map(req), &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "UpdateAdGroup httpclient.NewClient().Post error: %v", err)
		return nil, err
	}

	return &resp, nil
}

type UpdateAdGroupReq struct {
	AdvertiserID int64  `json:"advertiser_id"`           // 千川广告账户ID，必填
	CampaignID   int64  `json:"campaign_id"`             // 广告组ID，必填
	BudgetMode   string `json:"budget_mode,omitempty"`   // 预算类型，详见【附录-预算类型】，允许值：BUDGET_MODE_DAY 日预算、BUDGET_MODE_INFINITE 预算不限
	Budget       string `json:"budget,omitempty"`        // 广告组预算，最多支持两位小数，当budget_mode为BUDGET_MODE_DAY时必填，预算单次修改幅度不能低于100元，且日预算不少于300元
	CampaignName string `json:"campaign_name,omitempty"` // 广告组名称，长度为1-100个字符，其中1个中文字符算2位
}

type UpdateAdGroupRespData struct {
	CampaignID int64 `json:"campaign_id"`
}
