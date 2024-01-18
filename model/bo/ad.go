package bo

import (
	"github.com/JingruiLea/ad_boost/model/ttypes"
)

type Ad struct {
	AdCreateTime    string                `json:"ad_create_time"`
	AdID            int64                 `json:"ad_id"`
	AdModifyTime    string                `json:"ad_modify_time"`
	AwemeInfo       []*AwemeInfo          `json:"aweme_info"`
	CampaignId      int64                 `json:"campaign_id"`
	CampaignScene   ttypes.CampaignScene  `json:"campaign_scene"`
	DeliverySetting *DeliverySetting      `json:"delivery_setting"`
	LabAdType       ttypes.LabAdType      `json:"lab_ad_type"`
	MarketingGoal   ttypes.MarketingGoal  `json:"marketing_goal"`
	MarketingScene  ttypes.MarketingScene `json:"marketing_scene"`
	Name            string                `json:"name"`
	OptStatus       ttypes.OptStatus      `json:"opt_status"`
	ProductInfo     []interface{}         `json:"product_info"`
	Status          ttypes.AdStatus       `json:"status"`
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

const FulltimeSchedule = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"

type CreateAd struct {
	AdvertiserID                  int64                               `json:"advertiser_id"`
	MarketingGoal                 ttypes.MarketingGoal                `json:"marketing_goal"`
	CampaignScene                 ttypes.CampaignScene                `json:"campaign_scene"`
	MarketingScene                ttypes.MarketingScene               `json:"marketing_scene"`
	Name                          string                              `json:"name"`
	CampaignId                    int64                               `json:"campaign_id,omitempty"`    //千川广告组ID
	IsIntelligent                 int                                 `json:"is_intelligent,omitempty"` //是否开启智能投放, 0:不开启, 1:开启
	LabAdType                     ttypes.LabAdType                    `json:"lab_ad_type,omitempty"`
	AwemeID                       int64                               `json:"aweme_id,omitempty"`         //即将开播的抖音号ID
	ProductIds                    []int                               `json:"product_ids,omitempty"`      //品ID列表, 仅当营销目标为视频带货时必填
	DeliverySetting               *DeliverySetting                    `json:"delivery_setting,omitempty"` //投放设置（预算与出价）
	Audience                      *Audience                           `json:"audience,omitempty"`         //人群定向
	CreativeMaterialMode          ttypes.CreativeMaterialMode         `json:"creative_material_mode,omitempty"`
	FirstIndustryId               int                                 `json:"first_industry_id,omitempty"`
	SecondIndustryId              int                                 `json:"second_industry_id,omitempty"`
	ThirdIndustryId               int                                 `json:"third_industry_id,omitempty"`
	AdKeywords                    []string                            `json:"ad_keywords,omitempty"`
	IsHomepageHide                int                                 `json:"is_homepage_hide,omitempty"`
	CreativeList                  []*CreateAdCreative                 `json:"creative_list,omitempty"`
	ProgrammaticCreativeMediaList []*ttypes.ProgrammaticCreativeMedia `json:"programmatic_creative_media_list,omitempty"`
	ProgrammaticCreativeTitleList []*ProgrammaticCreativeTitle        `json:"programmatic_creative_title_list,omitempty"`
	ProgrammaticCreativeCard      *ProgrammaticCreativeCard           `json:"programmatic_creative_card,omitempty"`
	CreativeAutoGenerate          int                                 `json:"creative_auto_generate,omitempty"`
	DynamicCreative               int                                 `json:"dynamic_creative,omitempty"`
	Keywords                      []*ttypes.Keyword                   `json:"keywords,omitempty"`
	TrackUrl                      *ttypes.TrackUrl                    `json:"track_url,omitempty"`
}

// NewLiveCommonAd 创建直播通投策略
func NewLiveCommonAd(name string, awemeID, accountID, adGroupID int64) *CreateAd {
	return &CreateAd{
		AdvertiserID:   accountID,
		MarketingGoal:  ttypes.MarketingGoalLivePromGoods,
		CampaignScene:  ttypes.CampaignSceneDailySale,
		MarketingScene: ttypes.MarketingSceneFeed,
		Name:           name,
		CampaignId:     adGroupID,
		IsIntelligent:  0,
		LabAdType:      ttypes.LabAdTypeNotLabAd,
		AwemeID:        awemeID,
		ProductIds:     nil,
		DeliverySetting: &DeliverySetting{
			SmartBidType:     ttypes.SmartBidTypeSmartBidCustom,
			QCPXMode:         ttypes.QcpxModeOff,
			BudgetMode:       ttypes.BudgetModeDay,
			LiveScheduleType: ttypes.LiveScheduleTypeScheduleFromNow,
			ScheduleTime:     FulltimeSchedule,
		},
		Audience:             nil,
		CreativeMaterialMode: ttypes.CreativeMaterialModeCustom,
		FirstIndustryId:      0,
		SecondIndustryId:     0,
		ThirdIndustryId:      0,
		AdKeywords:           nil,
		IsHomepageHide:       0,
		CreativeList: []*CreateAdCreative{
			{
				ImageMode: ttypes.ImageModeAwemeLiveRoom,
			},
		},
		ProgrammaticCreativeMediaList: nil,
		ProgrammaticCreativeTitleList: nil,
		ProgrammaticCreativeCard:      nil,
		CreativeAutoGenerate:          0,
		DynamicCreative:               0,
		Keywords:                      nil,
		TrackUrl:                      nil,
	}
}

type CreateAdCreative struct {
	ImageMode             ttypes.ImageMode              `json:"image_mode"`
	VideoMaterial         *ttypes.VideoMaterial         `json:"video_material,omitempty"`
	ImageMaterial         *ttypes.ImageMaterial         `json:"image_material,omitempty"`
	TitleMaterial         *ProgrammaticCreativeTitle    `json:"title_material,omitempty"`
	PromotionCardMaterial *ttypes.PromotionCardMaterial `json:"promotion_card_material,omitempty"`
}

func (a *CreateAd) WithRoi(roiGoal int32) *CreateAd {
	a.DeliverySetting.ROIGoal = float64(roiGoal)
	a.DeliverySetting.ExternalAction = ttypes.ExternalActionAdConvertTypeLiveSuccessorderPay
	a.DeliverySetting.DeepExternalAction = ttypes.DeepExternalActionAdConvertTypeLivePayRoi
	a.DeliverySetting.DeepBidType = ttypes.DeepBidTypeMin
	return a
}

func (a *CreateAd) Copy() *CreateAd {
	return &CreateAd{
		AdvertiserID:                  a.AdvertiserID,
		MarketingGoal:                 a.MarketingGoal,
		CampaignScene:                 a.CampaignScene,
		MarketingScene:                a.MarketingScene,
		Name:                          a.Name,
		CampaignId:                    a.CampaignId,
		IsIntelligent:                 a.IsIntelligent,
		LabAdType:                     a.LabAdType,
		AwemeID:                       a.AwemeID,
		ProductIds:                    a.ProductIds,
		DeliverySetting:               a.DeliverySetting,
		Audience:                      a.Audience,
		CreativeMaterialMode:          a.CreativeMaterialMode,
		FirstIndustryId:               a.FirstIndustryId,
		SecondIndustryId:              a.SecondIndustryId,
		ThirdIndustryId:               a.ThirdIndustryId,
		AdKeywords:                    a.AdKeywords,
		IsHomepageHide:                a.IsHomepageHide,
		CreativeList:                  a.CreativeList,
		ProgrammaticCreativeMediaList: a.ProgrammaticCreativeMediaList,
		ProgrammaticCreativeTitleList: a.ProgrammaticCreativeTitleList,
		ProgrammaticCreativeCard:      a.ProgrammaticCreativeCard,
		CreativeAutoGenerate:          a.CreativeAutoGenerate,
		DynamicCreative:               a.DynamicCreative,
		Keywords:                      a.Keywords,
		TrackUrl:                      a.TrackUrl,
	}
}

func (a *CreateAd) WithBudget(budget float64) *CreateAd {
	a.DeliverySetting.Budget = budget
	return a
}

func (a *CreateAd) WithBid(bid float64) *CreateAd {
	a.DeliverySetting.CPABid = bid
	return a
}

func (a *CreateAd) WithFollowUser() *CreateAd {
	a.DeliverySetting.ExternalAction = ttypes.ExternalActionAdConvertTypeNewFollowAction
	return a
}
