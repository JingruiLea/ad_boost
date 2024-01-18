package ad

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/bo"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetAdListByStatus(ctx context.Context, accountID int64, status ttypes.AdStatus, adGroupID int64, filter *Filter) ([]*model.Ad, error) {
	if filter == nil {
		filter = &Filter{
			MarketingGoal: ttypes.MarketingGoalLivePromGoods,
			Status:        status,
			CampaignId:    adGroupID,
		}
	} else {
		filter.MarketingGoal = ttypes.MarketingGoalLivePromGoods
		filter.Status = status
		filter.CampaignId = adGroupID
	}
	resp, err := GetAdList(ctx, &GetAdListReq{
		AdvertiserId: accountID,
		Filtering:    filter,
		Page:         1,
		PageSize:     100,
	})
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdListByStauts GetAdList error: %v", err)
		return nil, err
	}
	var ret []*model.Ad
	for _, ad := range resp.List {
		var reciever model.Ad
		ret = append(ret, reciever.FromBO(ad))
	}
	return ret, nil
}

func GetAdList(ctx context.Context, req *GetAdListReq) (*GetAdListRespData, error) {
	var resp GetAdListRespData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserId, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/ad/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdList httpclient.NewClient().AdGet error: %v", err)
		return nil, err
	}
	return &resp, nil
}

type GetAdListReq struct {
	AdvertiserId     int64          `json:"advertiser_id"`
	RequestAwemeInfo ttypes.BoolInt `json:"request_aweme_info,omitempty"`
	Filtering        *Filter        `json:"filtering,omitempty"`
	Page             int            `json:"page,omitempty"` //允许值：10, 20, 50, 100, 200，默认值：10
	PageSize         int            `json:"page_size,omitempty"`
}

type Filter struct {
	IDs               []int64               `json:"ids,omitempty"`
	AdName            string                `json:"ad_name,omitempty"`
	Status            ttypes.AdStatus       `json:"status,omitempty"`
	CampaignScene     []CampaignSceneFilter `json:"campaign_scene,omitempty"`
	MarketingGoal     ttypes.MarketingGoal  `json:"marketing_goal"`
	MarketingScene    MarketingSceneFilter  `json:"marketing_scene,omitempty"`
	CampaignId        int64                 `json:"campaign_id,omitempty"`
	AdCreateStartDate string                `json:"ad_create_start_date,omitempty"`
	AdCreateEndDate   string                `json:"ad_create_end_date,omitempty"`
	AdModifyTime      string                `json:"ad_modify_time,omitempty"`
	AwemeId           int64                 `json:"aweme_id,omitempty"`
	AutoManageFilter  AutoManageFilter      `json:"auto_manage_filter,omitempty"`
}

type MarketingSceneFilter string

const (
	MarketingSceneFilterAll          MarketingSceneFilter = "ALL"
	MarketingSceneFilterFeed         MarketingSceneFilter = "FEED"
	MarketingSceneFilterSearch       MarketingSceneFilter = "SEARCH"
	MarketingSceneFilterShoppingMall MarketingSceneFilter = "SHOPPING_MALL"
)

type AutoManageFilter string

const (
	AutoManageFilterAllFilter  AutoManageFilter = "ALL"
	AutoManageFilterAutoManage AutoManageFilter = "AUTO_MANAGE"
	AutoManageFilterNormal     AutoManageFilter = "NORMAL"
)

type CampaignSceneFilter string

const (
	CampaignSceneFilterDailySale                 CampaignSceneFilter = "DAILY_SALE"
	CampaignSceneFilterNewCustomerTransformation CampaignSceneFilter = "NEW_CUSTOMER_TRANSFORMATION"
)

type GetAdListRespData struct {
	FailList []interface{}    `json:"fail_list"`
	List     []*bo.Ad         `json:"list"`
	PageInfo *ttypes.PageInfo `json:"page_info"`
}
