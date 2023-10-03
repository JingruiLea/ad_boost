package ad

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetAdListByStatus(ctx context.Context, accountID int64, status ttypes.AdStatus) ([]*model.Ad, error) {
	resp, err := GetAdList(ctx, &GetAdListReq{
		AdvertiserId: accountID,
		Filtering: &Filter{
			MarketingGoal: ttypes.MarketingGoalLivePromGoods,
			Status:        status,
		},
		Page:     1,
		PageSize: 100,
	})
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdListByStauts GetAdList error: %v", err)
		return nil, err
	}
	var ret []*model.Ad
	for _, ad := range resp.List {
		ret = append(ret, ad.ToModel())
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
	AdvertiserId     int64     `json:"advertiser_id"`
	RequestAwemeInfo AwemeInfo `json:"request_aweme_info,omitempty"`
	Filtering        *Filter   `json:"filtering,omitempty"`
	Page             int       `json:"page,omitempty"` //允许值：10, 20, 50, 100, 200，默认值：10
	PageSize         int       `json:"page_size,omitempty"`
}

type Filter struct {
	IDs               []int64               `json:"ids,omitempty"`
	AdName            string                `json:"ad_name,omitempty"`
	Status            ttypes.AdStatus       `json:"status,omitempty"`
	CampaignScene     []CampaignSceneFilter `json:"campaign_scene,omitempty"`
	MarketingGoal     ttypes.MarketingGoal  `json:"marketing_goal"`
	MarketingScene    MarketingSceneFilter  `json:"marketing_scene,omitempty"`
	CampaignId        int                   `json:"campaign_id,omitempty"`
	AdCreateStartDate string                `json:"ad_create_start_date,omitempty"`
	AdCreateEndDate   string                `json:"ad_create_end_date,omitempty"`
	AdModifyTime      string                `json:"ad_modify_time,omitempty"`
	AwemeId           int                   `json:"aweme_id,omitempty"`
	AutoManageFilter  AutoManageFilter      `json:"auto_manage_filter,omitempty"`
}

// 定义枚举类型
type AwemeInfo int

const (
	AwemeInfoNoInclude AwemeInfo = 0
	AwemeInfoInclude   AwemeInfo = 1
)

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
	List     []*Ad            `json:"list"`
	PageInfo *ttypes.PageInfo `json:"page_info"`
}

type Ad struct {
	AdCreateTime    string                `json:"ad_create_time"`
	AdID            int64                 `json:"ad_id"`
	AdModifyTime    string                `json:"ad_modify_time"`
	AwemeInfo       []interface{}         `json:"aweme_info"`
	CampaignId      int64                 `json:"campaign_id"`
	CampaignScene   string                `json:"campaign_scene"`
	DeliverySetting *DeliverySetting      `json:"delivery_setting"`
	LabAdType       ttypes.LabAdType      `json:"lab_ad_type"`
	MarketingGoal   ttypes.MarketingGoal  `json:"marketing_goal"`
	MarketingScene  ttypes.MarketingScene `json:"marketing_scene"`
	Name            string                `json:"name"`
	OptStatus       ttypes.OptStatus      `json:"opt_status"`
	ProductInfo     []interface{}         `json:"product_info"`
	Status          ttypes.AdStatus       `json:"status"`
}

func (a *Ad) ToModel() *model.Ad {
	ret := &model.Ad{
		AdCreateTime:    a.AdCreateTime,
		AdID:            a.AdID,
		AdModifyTime:    a.AdModifyTime,
		CampaignId:      a.CampaignId,
		CampaignScene:   a.CampaignScene,
		LabAdType:       a.LabAdType,
		MarketingGoal:   a.MarketingGoal,
		MarketingScene:  a.MarketingScene,
		Name:            a.Name,
		OptStatus:       a.OptStatus,
		Status:          a.Status,
		DeliverySetting: nil,
	}
	if a.DeliverySetting != nil {
		ret.DeliverySetting = &model.DeliverySetting{
			AdID:               a.AdID,
			Budget:             a.DeliverySetting.Budget,
			BudgetMode:         a.DeliverySetting.BudgetMode,
			DeepBidType:        a.DeliverySetting.DeepBidType,
			DeepExternalAction: a.DeliverySetting.DeepExternalAction,
			EndTime:            a.DeliverySetting.EndTime,
			ExternalAction:     a.DeliverySetting.ExternalAction,
			ProductNewOpen:     a.DeliverySetting.ProductNewOpen,
			RoiGoal:            a.DeliverySetting.RoiGoal,
			SmartBidType:       a.DeliverySetting.SmartBidType,
			StartTime:          a.DeliverySetting.StartTime,
			CpaBid:             a.DeliverySetting.CpaBid,
		}
	}
	return ret
}

type DeliverySetting struct {
	Budget             float64                   `json:"budget"`
	BudgetMode         ttypes.BudgetMode         `json:"budget_mode"`
	DeepBidType        ttypes.DeepBidType        `json:"deep_bid_type,omitempty"`
	DeepExternalAction ttypes.DeepExternalAction `json:"deep_external_action,omitempty"`
	EndTime            string                    `json:"end_time"`
	ExternalAction     ttypes.ExternalAction     `json:"external_action"`
	ProductNewOpen     bool                      `json:"product_new_open"`
	RoiGoal            float64                   `json:"roi_goal,omitempty"`
	SmartBidType       ttypes.SmartBidType       `json:"smart_bid_type"`
	StartTime          string                    `json:"start_time"`
	CpaBid             float64                   `json:"cpa_bid,omitempty"`
}
