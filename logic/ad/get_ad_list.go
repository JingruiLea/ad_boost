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
		AdCreateTime:   a.AdCreateTime,
		AdID:           a.AdID,
		AdModifyTime:   a.AdModifyTime,
		CampaignId:     a.CampaignId,
		CampaignScene:  a.CampaignScene,
		LabAdType:      a.LabAdType,
		MarketingGoal:  a.MarketingGoal,
		MarketingScene: a.MarketingScene,
		Name:           a.Name,
		OptStatus:      a.OptStatus,
		Status:         a.Status,
		DeliverySetting: model.DeliverySetting{
			SmartBidType:          a.DeliverySetting.SmartBidType,
			ExternalAction:        a.DeliverySetting.ExternalAction,
			DeepExternalAction:    a.DeliverySetting.DeepExternalAction,
			DeepBidType:           a.DeliverySetting.DeepBidType,
			ROIGoal:               a.DeliverySetting.ROIGoal,
			Budget:                a.DeliverySetting.Budget,
			ReviveBudget:          a.DeliverySetting.ReviveBudget,
			BudgetMode:            a.DeliverySetting.BudgetMode,
			CPABid:                a.DeliverySetting.CPABid,
			VideoScheduleType:     a.DeliverySetting.VideoScheduleType,
			LiveScheduleType:      a.DeliverySetting.LiveScheduleType,
			StartTime:             a.DeliverySetting.StartTime,
			EndTime:               a.DeliverySetting.EndTime,
			ScheduleTime:          a.DeliverySetting.ScheduleTime,
			ScheduleFixedRange:    a.DeliverySetting.ScheduleFixedRange,
			EnableAutoPause:       a.DeliverySetting.EnableAutoPause,
			AutoManageStrategyCmd: a.DeliverySetting.AutoManageStrategyCmd,
			EnableFollowMaterial:  a.DeliverySetting.EnableFollowMaterial,
			ProductNewOpen:        a.DeliverySetting.ProductNewOpen,
			QCPXMode:              a.DeliverySetting.QCPXMode,
			AllowQCPX:             a.DeliverySetting.AllowQCPX,
		},
	}
	return ret
}

type DeliverySetting struct {
	SmartBidType          ttypes.SmartBidType       `json:"smart_bid_type"`           // 投放场景（出价方式）
	ExternalAction        ttypes.ExternalAction     `json:"external_action"`          // 转化目标
	DeepExternalAction    ttypes.DeepExternalAction `json:"deep_external_action"`     // 深度转化目标
	DeepBidType           ttypes.DeepBidType        `json:"deep_bid_type"`            // 深度出价方式
	ROIGoal               float64                   `json:"roi_goal"`                 // 支付ROI目标
	Budget                float64                   `json:"budget"`                   // 预算
	ReviveBudget          float64                   `json:"revive_budget"`            // 复活预算
	BudgetMode            ttypes.BudgetMode         `json:"budget_mode"`              // 预算类型
	CPABid                float64                   `json:"cpa_bid"`                  // 转化出价
	VideoScheduleType     ttypes.VideoScheduleType  `json:"video_schedule_type"`      // 短视频投放日期选择方式
	LiveScheduleType      ttypes.LiveScheduleType   `json:"live_schedule_type"`       // 直播间投放时段选择方式
	StartTime             string                    `json:"start_time"`               // 投放开始时间
	EndTime               string                    `json:"end_time"`                 // 投放结束时间
	ScheduleTime          string                    `json:"schedule_time"`            // 投放时段
	ScheduleFixedRange    int                       `json:"schedule_fixed_range"`     // 固定投放时长
	EnableAutoPause       int                       `json:"enable_auto_pause"`        // 是否启用超成本自动暂停
	AutoManageStrategyCmd int                       `json:"auto_manage_strategy_cmd"` // 托管策略
	EnableFollowMaterial  int                       `json:"enable_follow_material"`   // 是否优质素材自动同步投放
	ProductNewOpen        bool                      `json:"product_new_open"`         // 是否开启新品加速
	QCPXMode              ttypes.QcpxMode           `json:"qcpx_mode"`                // 智能优惠券状态
	AllowQCPX             bool                      `json:"allow_qcpx"`               // 是否支持智能优惠券
}

func (a *Ad) IsRoi() bool {
	return a.MarketingGoal == ttypes.MarketingGoalLivePromGoods &&
		a.DeliverySetting.ExternalAction == ttypes.ExternalActionAdConvertTypeLiveSuccessorderPay &&
		a.DeliverySetting.DeepExternalAction == ttypes.DeepExternalActionAdConvertTypeLivePayRoi &&
		a.DeliverySetting.DeepBidType == ttypes.DeepBidTypeMin
}

func (a *Ad) IsCpa() bool {
	return a.MarketingGoal == ttypes.MarketingGoalLivePromGoods &&
		a.DeliverySetting.ExternalAction == ttypes.ExternalActionAdConvertTypeLiveSuccessorderPay &&
		a.DeliverySetting.DeepExternalAction == ""
}
