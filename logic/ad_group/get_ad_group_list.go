package ad_group

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetAdGroupList(ctx context.Context, req *GetAdGroupListReq) (*GetAdGroupListRespData, error) {
	req.Page = 1
	req.PageSize = 10
	req.AdvertiserID = 1748031128935424
	req.Filter.MarketingGoal = ttypes.MarketingGoalLivePromGoods

	var resp GetAdGroupListResp
	err := httpclient.NewClient().Get(ctx, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/campaign_list/get/", httpclient.CommonHeader, &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdGroupList httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	fmt.Printf("GetAdGroupList respMap: %s", utils.GetJsonStr(resp))
	if resp.Code != 0 || resp.Data == nil {
		return nil, fmt.Errorf("GetAdGroupList resp.Code != 0 || resp.Data == nil")
	}

	return resp.Data, nil
}

type GetAdGroupListReq struct {
	AdvertiserID int64   `json:"advertiser_id"`
	Page         int     `json:"page"`
	PageSize     int     `json:"page_size"`
	Filter       *Filter `json:"filter"`
}

type Filter struct {
	IDs            []int64               `json:"ids,omitempty"`
	Name           string                `json:"name,omitempty"`
	MarketingGoal  ttypes.MarketingGoal  `json:"marketing_goal"`
	MarketingScene ttypes.MarketingScene `json:"marketing_scene,omitempty"`
	Status         AdGroupStatus         `json:"status,omitempty"`
}

type AdGroupStatus string

const (
	AdGroupStatusAll     AdGroupStatus = "ALL"
	AdGroupStatusEnable  AdGroupStatus = "ENABLE"
	AdGroupStatusDisable AdGroupStatus = "DISABLE"
	AdGroupStatusDelete  AdGroupStatus = "DELETE"
)

type GetAdGroupListResp struct {
	ttypes.BaseResp
	Data    *GetAdGroupListRespData `json:"data"`
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
}

type GetAdGroupListRespData struct {
	List     []*Campaign      `json:"list"`
	PageInfo *ttypes.PageInfo `json:"page_info"`
}

type Campaign struct {
	BudgetMode     string  `json:"budget_mode"`
	CreateDate     string  `json:"create_date"`
	ID             int64   `json:"id"`
	MarketingGoal  string  `json:"marketing_goal"`
	MarketingScene string  `json:"marketing_scene"`
	Name           string  `json:"name"`
	Status         string  `json:"status"`
	Budget         float64 `json:"budget"`
}

func (c *Campaign) ToModel() *model.AdGroup {
	ret := &model.AdGroup{
		BudgetMode:     c.BudgetMode,
		CreateDate:     c.CreateDate,
		AdGroupID:      c.ID,
		MarketingGoal:  c.MarketingGoal,
		MarketingScene: c.MarketingScene,
		Name:           c.Name,
		Status:         c.Status,
		Budget:         c.Budget,
	}
	return ret
}
