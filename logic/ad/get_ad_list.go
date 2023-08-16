package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetAdList(ctx context.Context, req *GetAdListReq) (*GetAdListRespData, error) {
	var resp GetAdListResp
	err := httpclient.NewClient().Get(ctx, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/ad/get/", httpclient.CommonHeader, &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdList httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	if resp.Code != 0 || resp.Data == nil {
		return nil, fmt.Errorf("GetAdList resp.Code != 0 || resp.Data == nil")
	}
	return resp.Data, nil
}

type GetAdListReq struct {
	AdvertiserId     int64     `json:"advertiser_id"`
	RequestAwemeInfo AwemeInfo `json:"request_aweme_info,omitempty"`
	Filtering        *Filter   `json:"filtering,omitempty"`
	Page             int       `json:"page,omitempty"` //允许值：10, 20, 50, 100, 200，默认值：10
	PageSize         int       `json:"page_size,omitempty"`
}

type Filter struct {
	Ids               []int                 `json:"ids,omitempty"`
	AdName            string                `json:"ad_name,omitempty"`
	Status            AdStatus              `json:"status,omitempty"`
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

type AdStatus string

const (
	AdStatusDeliveryOk              AdStatus = "DELIVERY_OK"
	AdStatusAudit                   AdStatus = "AUDIT"
	AdStatusReaudit                 AdStatus = "REAUDIT"
	AdStatusDelete                  AdStatus = "DELETE"
	AdStatusDisable                 AdStatus = "DISABLE"
	AdStatusDraft                   AdStatus = "DRAFT"
	AdStatusTimeNoReach             AdStatus = "TIME_NO_REACH"
	AdStatusTimeDone                AdStatus = "TIME_DONE"
	AdStatusNoSchedule              AdStatus = "NO_SCHEDULE"
	AdStatusCreate                  AdStatus = "CREATE"
	AdStatusOfflineAudit            AdStatus = "OFFLINE_AUDIT"
	AdStatusOfflineBudget           AdStatus = "OFFLINE_BUDGET"
	AdStatusOfflineBalance          AdStatus = "OFFLINE_BALANCE"
	AdStatusPreOfflineBudget        AdStatus = "PRE_OFFLINE_BUDGET"
	AdStatusPreOnline               AdStatus = "PRE_ONLINE"
	AdStatusFrozen                  AdStatus = "FROZEN"
	AdStatusError                   AdStatus = "ERROR"
	AdStatusAuditStatusError        AdStatus = "AUDIT_STATUS_ERROR"
	AdStatusAdvertiserOfflineBudget AdStatus = "ADVERTISER_OFFLINE_BUDGET"
	AdStatusAdvertiserPreOffline    AdStatus = "ADVERTISER_PRE_OFFLINE_BUDGET"
	AdStatusExternalUrlDisable      AdStatus = "EXTERNAL_URL_DISABLE"
	AdStatusLiveRoomOff             AdStatus = "LIVE_ROOM_OFF"
	AdStatusCampaignDisable         AdStatus = "CAMPAIGN_DISABLE"
	AdStatusCampaignOfflineBudget   AdStatus = "CAMPAIGN_OFFLINE_BUDGET"
	AdStatusCampaignPreOffline      AdStatus = "CAMPAIGN_PREOFFLINE_BUDGET"
	AdStatusSystemDisable           AdStatus = "SYSTEM_DISABLE"
	AdStatusQuotaDisable            AdStatus = "QUOTA_DISABLE"
	AdStatusRoi2Disable             AdStatus = "ROI2_DISABLE"
)

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

type GetAdListResp struct {
	ttypes.BaseResp
	Data *GetAdListRespData `json:"data"`
}

type Ad struct {
	AdCreateTime    string           `json:"ad_create_time"`
	AdID            int64            `json:"ad_id"`
	AdModifyTime    string           `json:"ad_modify_time"`
	AwemeInfo       []interface{}    `json:"aweme_info"`
	CampaignId      int64            `json:"campaign_id"`
	CampaignScene   string           `json:"campaign_scene"`
	DeliverySetting *DeliverySetting `json:"delivery_setting"`
	LabAdType       string           `json:"lab_ad_type"`
	MarketingGoal   string           `json:"marketing_goal"`
	MarketingScene  string           `json:"marketing_scene"`
	Name            string           `json:"name"`
	OptStatus       string           `json:"opt_status"`
	ProductInfo     []interface{}    `json:"product_info"`
	Status          string           `json:"status"`
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
