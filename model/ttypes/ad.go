package ttypes

const FulltimeSchedule = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"

type Ad struct {
	AdvertiserID                  int64                        `json:"advertiser_id"`
	MarketingGoal                 MarketingGoal                `json:"marketing_goal"`
	CampaignScene                 CampaignScene                `json:"campaign_scene"`
	MarketingScene                MarketingScene               `json:"marketing_scene"`
	Name                          string                       `json:"name"`
	CampaignId                    int64                        `json:"campaign_id,omitempty"`    //千川广告组ID
	IsIntelligent                 int                          `json:"is_intelligent,omitempty"` //是否开启智能投放, 0:不开启, 1:开启
	LabAdType                     LabAdType                    `json:"lab_ad_type,omitempty"`
	AwemeID                       int64                        `json:"aweme_id,omitempty"`         //即将开播的抖音号ID
	ProductIds                    []int                        `json:"product_ids,omitempty"`      //品ID列表, 仅当营销目标为视频带货时必填
	DeliverySetting               *DeliverySetting             `json:"delivery_setting,omitempty"` //投放设置（预算与出价）
	Audience                      *Audience                    `json:"audience,omitempty"`         //人群定向
	CreativeMaterialMode          CreativeMaterialMode         `json:"creative_material_mode,omitempty"`
	FirstIndustryId               int                          `json:"first_industry_id,omitempty"`
	SecondIndustryId              int                          `json:"second_industry_id,omitempty"`
	ThirdIndustryId               int                          `json:"third_industry_id,omitempty"`
	AdKeywords                    []string                     `json:"ad_keywords,omitempty"`
	IsHomepageHide                int                          `json:"is_homepage_hide,omitempty"`
	CreativeList                  []*Creative                  `json:"creative_list,omitempty"`
	ProgrammaticCreativeMediaList []*ProgrammaticCreativeMedia `json:"programmatic_creative_media_list,omitempty"`
	ProgrammaticCreativeTitleList []*ProgrammaticCreativeTitle `json:"programmatic_creative_title_list,omitempty"`
	ProgrammaticCreativeCard      *ProgrammaticCreativeCard    `json:"programmatic_creative_card,omitempty"`
	CreativeAutoGenerate          int                          `json:"creative_auto_generate,omitempty"`
	DynamicCreative               int                          `json:"dynamic_creative,omitempty"`
	Keywords                      []*Keyword                   `json:"keywords,omitempty"`
	TrackUrl                      *TrackUrl                    `json:"track_url,omitempty"`
}

// NewLiveCommonAd 创建直播通投策略
func NewLiveCommonAd(name string, awemeID, adID, adGroupID int64) *Ad {
	return &Ad{
		AdvertiserID:   adID,
		MarketingGoal:  MarketingGoalLivePromGoods,
		CampaignScene:  CampaignSceneDailySale,
		MarketingScene: MarketingSceneFeed,
		Name:           name,
		CampaignId:     adGroupID,
		IsIntelligent:  0,
		LabAdType:      LabAdTypeNotLabAd,
		AwemeID:        awemeID,
		ProductIds:     nil,
		DeliverySetting: &DeliverySetting{
			SmartBidType:     SmartBidTypeSmartBidCustom,
			QcpxMode:         QcpxModeOff,
			BudgetMode:       BudgetModeDay,
			LiveScheduleType: LiveScheduleTypeScheduleFromNow,
			ScheduleTime:     FulltimeSchedule,
		},
		Audience:             nil,
		CreativeMaterialMode: CreativeMaterialModeCustom,
		FirstIndustryId:      0,
		SecondIndustryId:     0,
		ThirdIndustryId:      0,
		AdKeywords:           nil,
		IsHomepageHide:       0,
		CreativeList: []*Creative{
			{
				ImageMode:             ImageModeAwemeLiveRoom,
				VideoMaterial:         nil,
				ImageMaterial:         nil,
				TitleMaterial:         nil,
				PromotionCardMaterial: nil,
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

func (a *Ad) WithRoi(roiGoal int32) *Ad {
	a.DeliverySetting.RoiGoal = float32(roiGoal)
	a.DeliverySetting.ExternalAction = ExternalActionAdConvertTypeLiveSuccessorderPay
	a.DeliverySetting.DeepExternalAction = DeepExternalActionAdConvertTypeLivePayRoi
	a.DeliverySetting.DeepBidType = DeepBidTypeMin
	return a
}

func (a *Ad) Copy() *Ad {
	return &Ad{
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

func (a *Ad) WithBudget(budget float32) *Ad {
	a.DeliverySetting.Budget = budget
	return a
}

func (a *Ad) WithBid(bid float32) *Ad {
	a.DeliverySetting.CpaBid = bid
	return a
}

func (a *Ad) WithFollowUser() *Ad {
	a.DeliverySetting.ExternalAction = ExternalActionAdConvertTypeNewFollowAction
	return a
}

func NewAudience() *Audience {
	return &Audience{
		AudienceMode:           AudienceModeCustom,
		ExcludeLimitedRegion:   1,
		DistrictType:           false,
		District:               "",
		City:                   nil,
		LocationType:           "",
		Gender:                 "",
		Age:                    nil,
		AwemeFanBehaviors:      AwemeFanBehaviorsCommentedUser.Common(),
		AwemeFanBehaviorsDays:  AwemeFanBehaviorsDays60,
		AwemeFanCategories:     nil,
		AwemeFanAccounts:       nil,
		AutoExtendEnabled:      0,
		AutoExtendTargets:      nil,
		Platform:               nil,
		SmartInterestAction:    "",
		ActionScene:            nil,
		ActionDays:             0,
		ActionCategories:       nil,
		ActionWords:            nil,
		InterestCategories:     nil,
		InterestWords:          nil,
		Ac:                     nil,
		RetargetingTagsInclude: nil,
		RetargetingTagsExclude: nil,
		LivePlatformTags:       nil,
		NewCustomer:            "",
	}
}

func (a *Audience) WithYounger() *Audience {
	a.Age = []Age{
		AgeBetween1823, AgeBetween2430, AgeBetween3140,
	}
	return a
}

func (a *Audience) WithOlder() *Audience {
	a.Age = []Age{
		AgeBetween3140, AgeBetween4149, AgeAbove50,
	}
	return a
}

func (a *Audience) WithMiddle() *Audience {
	a.Age = []Age{
		AgeBetween2430, AgeBetween3140, AgeBetween4149,
	}
	return a
}

type Creative struct {
	ImageMode             ImageMode                  `json:"image_mode"`
	VideoMaterial         *VideoMaterial             `json:"video_material,omitempty"`
	ImageMaterial         *ImageMaterial             `json:"image_material,omitempty"`
	TitleMaterial         *ProgrammaticCreativeTitle `json:"title_material,omitempty"`
	PromotionCardMaterial *PromotionCardMaterial     `json:"promotion_card_material,omitempty"`
}

type VideoMaterial struct {
	VideoId      string `json:"video_id"`
	VideoCoverId string `json:"video_cover_id"`
	AwemeItemId  int    `json:"aweme_item_id"`
}

type ImageMaterial struct {
	ImageIds []string `json:"image_ids"`
}

type Audience struct {
	AudienceMode           AudienceMode          `json:"audience_mode"`
	OrientationId          int                   `json:"orientation_id"`
	ExcludeLimitedRegion   int                   `json:"exclude_limited_region"` //是否排除限运区域，0:不排除，1:排除
	DistrictType           bool                  `json:"district_type,omitempty"`
	District               string                `json:"district,omitempty"`
	City                   []int                 `json:"city,omitempty"`
	LocationType           LocationType          `json:"location_type,omitempty"`
	Gender                 Gender                `json:"gender"`
	Age                    []Age                 `json:"age"`
	AwemeFanBehaviors      []AwemeFanBehavior    `json:"aweme_fan_behaviors"`
	AwemeFanBehaviorsDays  AwemeFanBehaviorsDays `json:"aweme_fan_behaviors_days"`
	AwemeFanCategories     []int                 `json:"aweme_fan_categories"`
	AwemeFanAccounts       []int                 `json:"aweme_fan_accounts"`
	AutoExtendEnabled      int                   `json:"auto_extend_enabled"`
	AutoExtendTargets      []string              `json:"auto_extend_targets"`
	Platform               []string              `json:"platform,omitempty"`
	SmartInterestAction    string                `json:"smart_interest_action"`
	ActionScene            []string              `json:"action_scene"`
	ActionDays             int                   `json:"action_days"`
	ActionCategories       []int                 `json:"action_categories"`
	ActionWords            []int                 `json:"action_words"`
	InterestCategories     []int                 `json:"interest_categories"`
	InterestWords          []int                 `json:"interest_words"`
	Ac                     []string              `json:"ac"`
	RetargetingTagsInclude []int                 `json:"retargeting_tags_include"`
	RetargetingTagsExclude []int                 `json:"retargeting_tags_exclude"`
	LivePlatformTags       []string              `json:"live_platform_tags"`
	NewCustomer            string                `json:"new_customer"`
}

type PromotionCardMaterial struct {
	Title                   string   `json:"title"`
	SellingPoints           []string `json:"selling_points"`
	ImageId                 string   `json:"image_id"`
	ActionButton            string   `json:"action_button"`
	ButtonSmartOptimization int      `json:"button_smart_optimization"`
	ComponentId             int      `json:"component_id"`
}

type ProgrammaticCreativeMedia struct {
	ImageMode    string   `json:"image_mode"`
	VideoId      string   `json:"video_id"`
	VideoCoverId string   `json:"video_cover_id"`
	ImageIds     []string `json:"image_ids"`
	AwemeItemId  int      `json:"aweme_item_id"`
}

type ProgrammaticCreativeTitle struct {
	Title        string         `json:"title"`
	DynamicWords []*DynamicWord `json:"dynamic_words"`
	TitleType    string         `json:"title_type"`
}

type DynamicWord struct {
	WordId      int    `json:"word_id"`
	DictName    string `json:"dict_name"`
	DefaultWord string `json:"default_word"`
}

type Keyword struct {
	Word      string `json:"word"`
	MatchType string `json:"match_type"`
}

type TrackUrl struct {
	ShowTrackUrl               []string `json:"show_track_url"`
	ActionTrackUrl             []string `json:"action_track_url"`
	VideoPlayEffectiveTrackUrl []string `json:"video_play_effective_track_url"`
}

type ProgrammaticCreativeCard struct {
	PromotionCardTitle                   string   `json:"promotion_card_title"`
	PromotionCardSellingPoints           []string `json:"promotion_card_selling_points"`
	PromotionCardImageId                 string   `json:"promotion_card_image_id"`
	PromotionCardActionButton            string   `json:"promotion_card_action_button"`
	PromotionCardButtonSmartOptimization int      `json:"promotion_card_button_smart_optimization"`
}

type DeliverySetting struct {
	SmartBidType       SmartBidType       `json:"smart_bid_type"`
	QcpxMode           QcpxMode           `json:"qcpx_mode"`
	ExternalAction     ExternalAction     `json:"external_action"`
	DeepExternalAction DeepExternalAction `json:"deep_external_action,omitempty"`
	DeepBidType        DeepBidType        `json:"deep_bid_type,omitempty"`
	RoiGoal            float32            `json:"roi_goal,omitempty"` //支付ROI目标，最多支持两位小数，0.01～100
	//预算，最多支持两位小数
	//当预算模式为日预算时，预算范围是300 - 9999999.99；
	//当预算模式为总预算时，预算范围是max(300,投放天数x100) - 9999999.99
	//注意：托管计划仅支持日预算的预算范围
	Budget                float32           `json:"budget,omitempty"`
	BudgetMode            BudgetMode        `json:"budget_mode,omitempty"`
	CpaBid                float32           `json:"cpa_bid,omitempty"` //出价范围0.1-10000，2位小数
	VideoScheduleType     VideoScheduleType `json:"video_schedule_type,omitempty"`
	LiveScheduleType      LiveScheduleType  `json:"live_schedule_type,omitempty"`
	StartTime             string            `json:"start_time,omitempty"`
	EndTime               string            `json:"end_time,omitempty"`
	ScheduleTime          string            `json:"schedule_time,omitempty"`
	ScheduleFixedRange    int               `json:"schedule_fixed_range,omitempty"`
	EnableAutoPause       int               `json:"enable_auto_pause,omitempty"`        //是否开启超成本自动暂停，0:不开启，1:开启,注意：仅短视频+开启托管时，支持
	AutoManageStrategyCmd int               `json:"auto_manage_strategy_cmd,omitempty"` //托管策略，0 优先跑量 1 优先成本,注意：仅短视频+开启托管时，支持
	EnableFollowMaterial  int               `json:"enable_follow_material,omitempty"`   //是否优质素材自动同步投放 0 关闭 1 开启,注意：仅短视频+开启托管时，支持
}

type MarketingGoal string

const (
	// MarketingGoalLivePromGoods LIVE_PROM_GOODS 代表直播带货
	MarketingGoalLivePromGoods MarketingGoal = "LIVE_PROM_GOODS"
	// MarketingGoalVideoPromGoods VIDEO_PROM_GOODS 代表视频带货
	MarketingGoalVideoPromGoods MarketingGoal = "VIDEO_PROM_GOODS"
)

type CampaignScene string

const (
	// CampaignSceneDailySale DAILY_SALE 代表日常销售
	CampaignSceneDailySale CampaignScene = "DAILY_SALE"
)

type MarketingScene string

const (
	// MarketingSceneFeed FEED 信息流广告
	MarketingSceneFeed MarketingScene = "FEED"
	// MarketingSceneSearch SEARCH 搜索广告
	MarketingSceneSearch MarketingScene = "SEARCH"
)

// LabAdType
// 注意：
// 1. 短视频带货+通投 ，此时抖音号的bind_type只能为OFFICIAL或SELF。抖音号关系类型参考【附录-抖音号授权类型】
// 2.直播带货且当营销场景为搜索广告时，只能传LAB_AD
type LabAdType string

const (
	// LabAdTypeNotLabAd NOT_LAB_AD 自定义
	LabAdTypeNotLabAd LabAdType = "NOT_LAB_AD"
	// LabAdTypeLabAd LAB_AD 托管
	LabAdTypeLabAd LabAdType = "LAB_AD"
)

// CreativeMaterialMode 创意素材类型
// 若抖音号为“合作达人”类型，仅支持CUSTOM_CREATIVE
// 若抖音号为“官方/自运营”类型：
// 若为通投广告：托管仅支持PROGRAMMATIC_CREATIVE；自定义仅支持CUSTOM_CREATIVE
// 若为搜索广告，仅支持CUSTOM_CREATIVE
type CreativeMaterialMode string

const (
	// CreativeMaterialModeProgrammatic PROGRAMMATIC_CREATIVE 程序化创意
	CreativeMaterialModeProgrammatic CreativeMaterialMode = "PROGRAMMATIC_CREATIVE"
	// CreativeMaterialModeCustom CUSTOM_CREATIVE 自定义创意
	CreativeMaterialModeCustom CreativeMaterialMode = "CUSTOM_CREATIVE"
)

type ImageMode string

const (
	// ImageModeSquare 方图, 宽高比=1:1, 尺寸≥600*600, 文件大小≤10M, 大小10M以下, 图片格式支持jpg、jpeg、png、bmp、gif
	ImageModeSquare ImageMode = "SQUARE"
	// ImageModeSmall 小图, 宽高比1.52, 大小1.5M以下, 下限：456 & 300, 上限：1368 & 900
	ImageModeSmall ImageMode = "SMALL"
	// ImageModeLarge 大图, 横版大图宽高比1.78, 大小1.5M以下, 下限：1280 & 720, 上限：2560 & 1440
	ImageModeLarge ImageMode = "LARGE"
	// ImageModeVideoLarge 横版视频, 封面图宽高比1.78（下限：1280 & 720，上限：2560 & 1440））,视频资源支持mp4、mpeg、3gp、avi文件格式，宽高比16:9，大小不超过1000M
	ImageModeVideoLarge ImageMode = "VIDEO_LARGE"
	// ImageModeVideoVertical 竖版视频, 封面图宽高比0.56（9:16），下限：720 & 1280，上限：1440 & 2560，视频资源支持mp4、mpeg、3gp、avi文件格式，大小不超过100M
	ImageModeVideoVertical ImageMode = "VIDEO_VERTICAL"
	// ImageModeLargeVertical 大图竖图, 宽高比0.56, 大小1.5M以下, 下限：720 & 1280, 上限：1440 & 2560
	ImageModeLargeVertical ImageMode = "LARGE_VERTICAL"
	// ImageModeUnionSplash 穿山甲开屏图片, 仅限穿山甲开屏使用, 下限：1080 & 1920, 上限：2160 & 3840, 比例0.56
	ImageModeUnionSplash ImageMode = "UNION_SPLASH"
	// ImageModeAwemeLiveRoom 直播画面类型
	ImageModeAwemeLiveRoom ImageMode = "AWEME_LIVE_ROOM"
)

// SmartBidType
// 注意：
// 1.当“广告类型”为“搜索广告”时，smart_bid_type只支持SMART_BID_CUSTOM
// 2.托管计划仅支持SMART_BID_CUSTOM 控成本投放
type SmartBidType string

const (
	// SmartBidTypeSmartBidCustom SMART_BID_CUSTOM 控成本投放，控制成本，尽量消耗完预算
	SmartBidTypeSmartBidCustom SmartBidType = "SMART_BID_CUSTOM"
	// SmartBidTypeSmartBidConservative SMART_BID_CONSERVATIVE 放量投放，接受成本上浮，尽量消耗更多预算
	SmartBidTypeSmartBidConservative SmartBidType = "SMART_BID_CONSERVATIVE"
)

// ExternalAction 定义转化目标
type ExternalAction string

const (
	// ExternalActionAdConvertTypeShopping 商品购买
	ExternalActionAdConvertTypeShopping ExternalAction = "AD_CONVERT_TYPE_SHOPPING"
	// ExternalActionAdConvertTypeQCFollowAction 粉丝提升
	ExternalActionAdConvertTypeQCFollowAction ExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	// ExternalActionAdConvertTypeQCMustBuy 点赞评论
	ExternalActionAdConvertTypeQCMustBuy ExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	// ExternalActionAdConvertTypeLiveEnterAction 进入直播间
	ExternalActionAdConvertTypeLiveEnterAction ExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	// ExternalActionAdConvertTypeLiveClickProductAction 直播间商品点击
	ExternalActionAdConvertTypeLiveClickProductAction ExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	// ExternalActionAdConvertTypeLiveSuccessorderAction 直播间下单
	ExternalActionAdConvertTypeLiveSuccessorderAction ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	// ExternalActionAdConvertTypeNewFollowAction 直播间粉丝提升
	ExternalActionAdConvertTypeNewFollowAction ExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	// ExternalActionAdConvertTypeLiveCommentAction 直播间评论
	ExternalActionAdConvertTypeLiveCommentAction ExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	// ExternalActionAdConvertTypeLiveSuccessorderPay 直播间成交
	ExternalActionAdConvertTypeLiveSuccessorderPay ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	// ExternalActionAdConvertTypeLiveSuccessorderSettle 直播间结算
	ExternalActionAdConvertTypeLiveSuccessorderSettle ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE"
	// ExternalActionAdConvertTypeLiveSuccessorderPay7Days 直播间成交-7日总成交
	ExternalActionAdConvertTypeLiveSuccessorderPay7Days ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS"
)

// DeepExternalAction
// 注意：
// 1. 当 smart_bid_type 为SMART_BID_CONSERVATIVE&SMART_BID_CUSTOM时，不支持
// 2. 推直播间+搜索广告时，入参AD_CONVERT_TYPE_LIVE_PAY_ROI即代表设置优化目标为“支付ROI”，此时不允许入参deep_bid_type
type DeepExternalAction string

const (
	// DeepExternalActionAdConvertTypeLivePayRoi 支付ROI
	DeepExternalActionAdConvertTypeLivePayRoi DeepExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
)

// DeepBidType
// 注意：
// 1.仅当深度转化目标为AD_CONVERT_TYPE_LIVE_PAY_ROI 时，必填；否则，填入也会报错
// 2.推直播间+搜索广告时，不允许传入，否则报错
// 3.在非搜索场景下，仅支持MIN；在推商品-日常销售-搜索，仅支持PAY_ROI
type DeepBidType string

const (
	// DeepBidTypeMin MIN 最小化成本
	DeepBidTypeMin DeepBidType = "MIN"
	// DeepBidTypePayRoi PAY_ROI 支付ROI
	DeepBidTypePayRoi DeepBidType = "PAY_ROI"
)

// BudgetMode 预算类型
// 注意： 托管计划仅支持BUDGET_MODE_DAY 日预算
type BudgetMode string

const (
	// BudgetModeDay 日预算
	BudgetModeDay BudgetMode = "BUDGET_MODE_DAY"
	// BudgetModeTotal 总预算
	BudgetModeTotal BudgetMode = "BUDGET_MODE_TOTAL"
)

type VideoScheduleType string

const (
	// VideoScheduleTypeScheduleFromNow 从今天起长期投放
	VideoScheduleTypeScheduleFromNow VideoScheduleType = "SCHEDULE_FROM_NOW"
	// VideoScheduleTypeScheduleStartEnd 设置开始和结束日期
	VideoScheduleTypeScheduleStartEnd VideoScheduleType = "SCHEDULE_START_END"
)

type LiveScheduleType string

const (
	// LiveScheduleTypeScheduleFromNow 从今天起长期投放
	LiveScheduleTypeScheduleFromNow LiveScheduleType = "SCHEDULE_FROM_NOW"
	// LiveScheduleTypeScheduleStartEnd 设置开始和结束日期
	LiveScheduleTypeScheduleStartEnd LiveScheduleType = "SCHEDULE_START_END"
	// LiveScheduleTypeScheduleTimeFixedrange 固定时长
	LiveScheduleTypeScheduleTimeFixedrange LiveScheduleType = "SCHEDULE_TIME_FIXEDRANGE"
)

// QcpxMode 是否开启智能优惠券，允许值
type QcpxMode string

const (
	// QcpxModeOn 启用
	QcpxModeOn QcpxMode = "QCPX_MODE_ON"
	// QcpxModeOff 不启用
	QcpxModeOff QcpxMode = "QCPX_MODE_OFF"
)

type AudienceMode string

const (
	// AudienceModeCustom CUSTOM 自定义
	AudienceModeCustom AudienceMode = "CUSTOM"
	// AudienceModeNone NONE 不限
	AudienceModeNone AudienceMode = "NONE"
	// AudienceModeOrientationPackage ORIENTATION_PACKAGE 已有定向包
	AudienceModeOrientationPackage AudienceMode = "ORIENTATION_PACKAGE"
	// AudienceModeAutoOrientation AUTO_ORIENTATION 智能定向, 仅直播间成交时支持
	AudienceModeAutoOrientation AudienceMode = "AUTO_ORIENTATION"
)

type DistrictType string

const (
	// DistrictTypeCity CITY 省市
	DistrictTypeCity DistrictType = "CITY"
	// DistrictTypeCounty COUNTY 区县
	DistrictTypeCounty DistrictType = "COUNTY"
	// DistrictTypeNone NONE 不限，默认值
	DistrictTypeNone DistrictType = "NONE"
)

type LocationType string

const (
	// LocationTypeCurrent CURRENT 正在该地区的用户
	LocationTypeCurrent LocationType = "CURRENT"
	// LocationTypeHome HOME 居住在该地区的用户
	LocationTypeHome LocationType = "HOME"
	// LocationTypeTravel TRAVEL 到该地区旅行的用户
	LocationTypeTravel LocationType = "TRAVEL"
	// LocationTypeAll ALL 该地区内的所有用户
	LocationTypeAll LocationType = "ALL"
)

//

type Gender string

const (
	// GenderFemale 女性
	GenderFemale Gender = "GENDER_FEMALE"
	// GenderMale 男性
	GenderMale Gender = "GENDER_MALE"
	// GenderNone 不限
	GenderNone Gender = "NONE"
)

type Age string

const (
	// AgeBetween1823 18-23岁
	AgeBetween1823 Age = "AGE_BETWEEN_18_23"
	// AgeBetween2430 24-30岁
	AgeBetween2430 Age = "AGE_BETWEEN_24_30"
	// AgeBetween3140 31-40岁
	AgeBetween3140 Age = "AGE_BETWEEN_31_40"
	// AgeBetween4149 41-49岁
	AgeBetween4149 Age = "AGE_BETWEEN_41_49"
	// AgeAbove50 50岁以上
	AgeAbove50 Age = "AGE_ABOVE_50"
)

type AwemeFanBehavior string

func (*AwemeFanBehavior) All() []AwemeFanBehavior {
	return []AwemeFanBehavior{
		AwemeFanBehaviorsFollowedUser,
		AwemeFanBehaviorsCommentedUser,
		AwemeFanBehaviorsLikedUser,
		AwemeFanBehaviorsSharedUser,
		AwemeFanBehaviorsLiveWatch,
		AwemeFanBehaviorsLiveEffectiveWatch,
		AwemeFanBehaviorsLiveComment,
		AwemeFanBehaviorsLiveExceptional,
		AwemeFanBehaviorsLiveGoodsClick,
		AwemeFanBehaviorsLiveGoodsOrder,
		AwemeFanBehaviorsGoodsCartsClick,
		AwemeFanBehaviorsGoodsCartsOrder,
	}
}

func (*AwemeFanBehavior) Video() []AwemeFanBehavior {
	return []AwemeFanBehavior{
		AwemeFanBehaviorsFollowedUser,
		AwemeFanBehaviorsCommentedUser,
		AwemeFanBehaviorsLikedUser,
		AwemeFanBehaviorsSharedUser,
	}
}

func (*AwemeFanBehavior) Live() []AwemeFanBehavior {
	return []AwemeFanBehavior{
		AwemeFanBehaviorsLiveWatch,
		AwemeFanBehaviorsLiveEffectiveWatch,
		AwemeFanBehaviorsLiveComment,
		AwemeFanBehaviorsLiveExceptional,
		AwemeFanBehaviorsLiveGoodsClick,
		AwemeFanBehaviorsLiveGoodsOrder,
	}
}

func (*AwemeFanBehavior) Goods() []AwemeFanBehavior {
	return []AwemeFanBehavior{
		AwemeFanBehaviorsLiveGoodsClick,
		AwemeFanBehaviorsLiveGoodsOrder,
		AwemeFanBehaviorsGoodsCartsClick,
		AwemeFanBehaviorsGoodsCartsOrder,
	}
}

func (AwemeFanBehavior) Common() []AwemeFanBehavior {
	return []AwemeFanBehavior{
		AwemeFanBehaviorsLiveEffectiveWatch,
		AwemeFanBehaviorsLiveComment,
		AwemeFanBehaviorsLiveExceptional,
		AwemeFanBehaviorsLiveGoodsClick,
		AwemeFanBehaviorsLiveGoodsOrder,
		AwemeFanBehaviorsGoodsCartsClick,
		AwemeFanBehaviorsGoodsCartsOrder,
	}
}

const (
	// AwemeFanBehaviorsFollowedUser 视频互动-关注
	AwemeFanBehaviorsFollowedUser AwemeFanBehavior = "FOLLOWED_USER"
	// AwemeFanBehaviorsCommentedUser 视频互动-评论
	AwemeFanBehaviorsCommentedUser AwemeFanBehavior = "COMMENTED_USER"
	// AwemeFanBehaviorsLikedUser 视频互动-点赞
	AwemeFanBehaviorsLikedUser AwemeFanBehavior = "LIKED_USER"
	// AwemeFanBehaviorsSharedUser 视频互动-分享
	AwemeFanBehaviorsSharedUser AwemeFanBehavior = "SHARED_USER"
	// AwemeFanBehaviorsLiveWatch 直播互动-观看
	AwemeFanBehaviorsLiveWatch AwemeFanBehavior = "LIVE_WATCH"
	// AwemeFanBehaviorsLiveEffectiveWatch 直播互动-有效观看
	AwemeFanBehaviorsLiveEffectiveWatch AwemeFanBehavior = "LIVE_EFFECTIVE_WATCH"
	// AwemeFanBehaviorsLiveComment 直播互动-直播评论
	AwemeFanBehaviorsLiveComment AwemeFanBehavior = "LIVE_COMMENT"
	// AwemeFanBehaviorsLiveExceptional 直播互动-打赏
	AwemeFanBehaviorsLiveExceptional AwemeFanBehavior = "LIVE_EXCEPTIONAL"
	// AwemeFanBehaviorsLiveGoodsClick 直播互动-商品点击
	AwemeFanBehaviorsLiveGoodsClick AwemeFanBehavior = "LIVE_GOODS_CLICK"
	// AwemeFanBehaviorsLiveGoodsOrder 直播互动-商品下单
	AwemeFanBehaviorsLiveGoodsOrder AwemeFanBehavior = "LIVE_GOODS_ORDER"
	// AwemeFanBehaviorsGoodsCartsClick 商品互动-购物车点击
	AwemeFanBehaviorsGoodsCartsClick AwemeFanBehavior = "GOODS_CARTS_CLICK"
	// AwemeFanBehaviorsGoodsCartsOrder 商品互动-购物车下单
	AwemeFanBehaviorsGoodsCartsOrder AwemeFanBehavior = "GOODS_CARTS_ORDER"
)

type AwemeFanBehaviorsDays string

const (
	// AwemeFanBehaviorsDays15 近15天
	AwemeFanBehaviorsDays15 AwemeFanBehaviorsDays = "DAYS_15"
	// AwemeFanBehaviorsDays30 近30天
	AwemeFanBehaviorsDays30 AwemeFanBehaviorsDays = "DAYS_30"
	// AwemeFanBehaviorsDays60 近60天
	AwemeFanBehaviorsDays60 AwemeFanBehaviorsDays = "DAYS_60"
)

// Platform 不传值为全选
type Platform string

const (
	// PlatformAndroid ANDROID 安卓
	PlatformAndroid Platform = "ANDROID"
	// PlatformIos IOS 苹果
	PlatformIos Platform = "IOS"
)

// SmartInterestAction RECOMMEND系统推荐，CUSTOM 自定义；不传值则为不限制
type SmartInterestAction string

const (
	// SmartInterestActionRecommend RECOMMEND 系统推荐
	SmartInterestActionRecommend SmartInterestAction = "RECOMMEND"
	// SmartInterestActionCustom CUSTOM 自定义
	SmartInterestActionCustom SmartInterestAction = "CUSTOM"
)

// 行为场景，详见【附录-行为场景】
// smart_interest_actionCUSTOM时有效，允许值：
// E-COMMERCE 电商互动行为、NEWS 资讯互动行为、APP APP推广互动行为
// 注意：仅以下两种情况支持：短视频带货+通投广告+非托管计划支持、直播带货+通投广告支持
type ActionScene string

const (
	// ActionSceneECommerce E-COMMERCE 电商互动行为
	ActionSceneECommerce ActionScene = "E-COMMERCE"
	// ActionSceneNews NEWS 资讯互动行为
	ActionSceneNews ActionScene = "NEWS"
	// ActionSceneApp APP APP推广互动行为
	ActionSceneApp ActionScene = "APP"
)

// ActionDays 用户发生行为天数
// 当 smart_interest_action 传 CUSTOM 时有效
// 注意：仅以下两种情况支持：短视频带货+通投广告+非托管计划支持、直播带货+通投广告支持
type ActionDays int

const (
	ActionDays7   ActionDays = 7
	ActionDays15  ActionDays = 15
	ActionDays30  ActionDays = 30
	ActionDays60  ActionDays = 60
	ActionDays90  ActionDays = 90
	ActionDays180 ActionDays = 180
	ActionDays365 ActionDays = 365
)

type AC string

func (*AC) All() []AC {
	return []AC{
		ACWifi,
		AC2G,
		AC3G,
		AC4G,
	}
}

const (
	// ACWifi WIFI
	ACWifi AC = "WIFI"
	// AC2G 2G
	AC2G AC = "2G"
	// AC3G 3G
	AC3G AC = "3G"
	// AC4G 4G
	AC4G AC = "4G"
)

// LivePlatformTags 注意：仅通投广告支持
type LivePlatformTags string

const (
	// LivePlatformTagsLargeFanscount LARGE_FANSCOUNT 高关注人群
	LivePlatformTagsLargeFanscount LivePlatformTags = "LARGE_FANSCOUNT"
	// LivePlatformTagsAbnormalActive ABNORMAL_ACTIVE 高活跃人群
	LivePlatformTagsAbnormalActive LivePlatformTags = "ABNORMAL_ACTIVE"
	// LivePlatformTagsAwemeFans AWEME_FANS 抖音号粉丝
	LivePlatformTagsAwemeFans LivePlatformTags = "AWEME_FANS"
)

// NewCustomer 注意：仅通投×自定义时，支持
type NewCustomer string

const (
	// NewCustomerNoBuy 店铺未购
	NewCustomerNoBuy NewCustomer = "NO_BUY"
	// NewCustomerNone 不限
	NewCustomerNone NewCustomer = "NONE"
)

type OptStatus string

const (
	OptStatusEnable  OptStatus = "ENABLE"
	OptStatusDisable OptStatus = "DISABLE"
	OptStatusDelete  OptStatus = "DELETE"
	OptStatusRevive  OptStatus = "REVIVE"
)

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
