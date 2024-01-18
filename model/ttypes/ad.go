package ttypes

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
	OrientationId          int64                 `json:"orientation_id,omitempty"`
	ExcludeLimitedRegion   int                   `json:"exclude_limited_region"` //是否排除限运区域，0:不排除，1:排除
	DistrictType           bool                  `json:"district_type,omitempty"`
	District               string                `json:"district,omitempty"`
	City                   []int                 `json:"city,omitempty"`
	LocationType           LocationType          `json:"location_type,omitempty"`
	Gender                 Gender                `json:"gender,omitempty"`
	Age                    []Age                 `json:"age,omitempty"`
	AwemeFanBehaviors      []AwemeFanBehavior    `json:"aweme_fan_behaviors,omitempty"`
	AwemeFanBehaviorsDays  AwemeFanBehaviorsDays `json:"aweme_fan_behaviors_days,omitempty"`
	AwemeFanCategories     []int                 `json:"aweme_fan_categories,omitempty"`
	AwemeFanAccounts       []int                 `json:"aweme_fan_accounts,omitempty"`
	AutoExtendEnabled      int                   `json:"auto_extend_enabled"`
	AutoExtendTargets      []string              `json:"auto_extend_targets,omitempty"`
	Platform               []string              `json:"platform,omitempty,omitempty"`
	SmartInterestAction    string                `json:"smart_interest_action,omitempty"`
	ActionScene            []ActionScene         `json:"action_scene,omitempty"`
	ActionDays             ActionDays            `json:"action_days,omitempty"`
	ActionCategories       []int                 `json:"action_categories,omitempty"`
	ActionWords            []int                 `json:"action_words,omitempty"`
	InterestCategories     []int                 `json:"interest_categories,omitempty"`
	InterestWords          []int                 `json:"interest_words,omitempty"`
	Ac                     []string              `json:"ac,omitempty"`
	RetargetingTagsInclude []int                 `json:"retargeting_tags_include,omitempty"`
	RetargetingTagsExclude []int                 `json:"retargeting_tags_exclude,omitempty"`
	LivePlatformTags       []string              `json:"live_platform_tags,omitempty"`
	NewCustomer            string                `json:"new_customer,omitempty"`
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

type MarketingGoal string

const (
	// MarketingGoalLiveAll ALL 只能查询时使用
	MarketingGoalLiveAll MarketingGoal = "ALL"
	// MarketingGoalLivePromGoods LIVE_PROM_GOODS 代表直播带货
	MarketingGoalLivePromGoods MarketingGoal = "LIVE_PROM_GOODS"
	// MarketingGoalVideoPromGoods VIDEO_PROM_GOODS 代表视频带货
	MarketingGoalVideoPromGoods MarketingGoal = "VIDEO_PROM_GOODS"
)

func (m MarketingGoal) String() string {
	switch m {
	case MarketingGoalLiveAll:
		return "全部"
	case MarketingGoalLivePromGoods:
		return "直播带货"
	case MarketingGoalVideoPromGoods:
		return "视频带货"
	default:
		return "未知营销目标"
	}
}

type CampaignScene string

const (
	// CampaignSceneDailySale DAILY_SALE 代表日常销售
	CampaignSceneDailySale CampaignScene = "DAILY_SALE"
)

func (c CampaignScene) String() string {
	switch c {
	case CampaignSceneDailySale:
		return "日常销售"
	default:
		return "未知营销场景"
	}
}

type MarketingScene string

const (
	// MarketingSceneFeed FEED 信息流广告
	MarketingSceneFeed MarketingScene = "FEED"
	// MarketingSceneSearch SEARCH 搜索广告
	MarketingSceneSearch MarketingScene = "SEARCH"
)

func (m MarketingScene) String() string {
	switch m {
	case MarketingSceneFeed:
		return "信息流广告"
	case MarketingSceneSearch:
		return "搜索广告"
	default:
		return "未知营销场景"
	}
}

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

func (l LabAdType) String() string {
	switch l {
	case LabAdTypeNotLabAd:
		return "自定义"
	case LabAdTypeLabAd:
		return "托管"
	default:
		return "未知"
	}
}

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

func (c CreativeMaterialMode) String() string {
	switch c {
	case CreativeMaterialModeProgrammatic:
		return "程序化创意"
	case CreativeMaterialModeCustom:
		return "自定义创意"
	default:
		return "未知"
	}
}

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

func (s SmartBidType) String() string {
	switch s {
	case SmartBidTypeSmartBidCustom:
		return "控成本投放"
	case SmartBidTypeSmartBidConservative:
		return "放量投放"
	default:
		return "未知"
	}
}

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

func (e ExternalAction) String() string {
	switch e {
	case ExternalActionAdConvertTypeShopping:
		return "商品购买"
	case ExternalActionAdConvertTypeQCFollowAction:
		return "粉丝提升"
	case ExternalActionAdConvertTypeQCMustBuy:
		return "点赞评论"
	case ExternalActionAdConvertTypeLiveEnterAction:
		return "进入直播间"
	case ExternalActionAdConvertTypeLiveClickProductAction:
		return "直播间商品点击"
	case ExternalActionAdConvertTypeLiveSuccessorderAction:
		return "直播间下单"
	case ExternalActionAdConvertTypeNewFollowAction:
		return "直播间粉丝提升"
	case ExternalActionAdConvertTypeLiveCommentAction:
		return "直播间评论"
	case ExternalActionAdConvertTypeLiveSuccessorderPay:
		return "直播间成交"
	case ExternalActionAdConvertTypeLiveSuccessorderSettle:
		return "直播间结算"
	case ExternalActionAdConvertTypeLiveSuccessorderPay7Days:
		return "直播间成交-7日总成交"
	default:
		return "未知"
	}
}

// DeepExternalAction
// 注意：
// 1. 当 smart_bid_type 为SMART_BID_CONSERVATIVE&SMART_BID_CUSTOM时，不支持
// 2. 推直播间+搜索广告时，入参AD_CONVERT_TYPE_LIVE_PAY_ROI即代表设置优化目标为“支付ROI”，此时不允许入参deep_bid_type
type DeepExternalAction string

const (
	// DeepExternalActionAdConvertTypeLivePayRoi 支付ROI
	DeepExternalActionAdConvertTypeLivePayRoi DeepExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
)

func (d DeepExternalAction) String() string {
	switch d {
	case DeepExternalActionAdConvertTypeLivePayRoi:
		return "支付ROI"
	default:
		return "未知"
	}
}

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

func (d DeepBidType) String() string {
	switch d {
	case DeepBidTypeMin:
		return "最小化成本"
	case DeepBidTypePayRoi:
		return "支付ROI"
	default:
		return "未知"
	}
}

// BudgetMode 预算类型
// 注意： 托管计划仅支持BUDGET_MODE_DAY 日预算
type BudgetMode string

const (
	// BudgetModeDay 日预算
	BudgetModeDay BudgetMode = "BUDGET_MODE_DAY"
	// BudgetModeTotal 总预算
	BudgetModeTotal BudgetMode = "BUDGET_MODE_TOTAL"
	//BUDGET_MODE_INFINITE 不限预算
	BudgetModeInfinite BudgetMode = "BUDGET_MODE_INFINITE"
)

func (b BudgetMode) String() string {
	switch b {
	case BudgetModeDay:
		return "日预算"
	case BudgetModeTotal:
		return "总预算"
	case BudgetModeInfinite:
		return "不限预算"
	default:
		return "未知预算类型"
	}
}

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

type District string

const (
	// DistrictCity CITY 省市
	DistrictCity District = "CITY"
	// DistrictCounty COUNTY 区县
	DistrictCounty District = "COUNTY"
	// DistrictNone NONE 不限，默认值
	DistrictNone District = "NONE"
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

func (o OptStatus) String() string {
	switch o {
	case OptStatusEnable:
		return "启用"
	case OptStatusDisable:
		return "暂停"
	case OptStatusDelete:
		return "删除"
	case OptStatusRevive:
		return "恢复"
	default:
		return "未知"
	}
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
	AdStatusSystemDisable           AdStatus = "SYSTEM_DISABLE" //系统暂停，因低效计划被系统自动暂停
	AdStatusQuotaDisable            AdStatus = "QUOTA_DISABLE"
	AdStatusRoi2Disable             AdStatus = "ROI2_DISABLE" //因该计划关联的抖音号开启全域推广，因此本计划被系统暂停
)

func (a AdStatus) String() string {
	switch a {
	case AdStatusDeliveryOk:
		return "投放中"
	case AdStatusAudit:
		return "审核中"
	case AdStatusReaudit:
		return "审核中"
	case AdStatusDelete:
		return "已删除"
	case AdStatusDisable:
		return "暂停"
	case AdStatusDraft:
		return "草稿"
	case AdStatusTimeNoReach:
		return "未到投放时间"
	case AdStatusTimeDone:
		return "已过投放时间"
	case AdStatusNoSchedule:
		return "未到投放时间"
	case AdStatusCreate:
		return "待审核"
	case AdStatusOfflineAudit:
		return "审核中"
	case AdStatusOfflineBudget:
		return "预算不足"
	case AdStatusOfflineBalance:
		return "余额不足"
	case AdStatusPreOfflineBudget:
		return "预算不足"
	case AdStatusPreOnline:
		return "待审核"
	case AdStatusFrozen:
		return "已冻结"
	case AdStatusError:
		return "审核失败"
	case AdStatusAuditStatusError:
		return "审核失败"
	case AdStatusAdvertiserOfflineBudget:
		return "预算不足"
	case AdStatusAdvertiserPreOffline:
		return "预算不足"
	case AdStatusExternalUrlDisable:
		return "外部链接不可用"
	case AdStatusLiveRoomOff:
		return "直播间已关闭"
	case AdStatusCampaignDisable:
		return "计划已暂停"
	case AdStatusCampaignOfflineBudget:
		return "计划预算不足"
	case AdStatusCampaignPreOffline:
		return "计划预算不足"
	case AdStatusSystemDisable:
		return "系统暂停"
	case AdStatusQuotaDisable:
		return "额度暂停"
	case AdStatusRoi2Disable:
		return "ROI2暂停"
	default:
		return "未知"
	}
}

type ECPType string

// 账户类型，可选值:
// SHOP: 商家
// SHOP_STAR: 商家达人
// COMMON_STAR: 普通达人
// AGENT: 百应机构
const (
	ECPTypeShop       ECPType = "SHOP"
	ECPTypeShopStar   ECPType = "SHOP_STAR"
	ECPTypeCommonStar ECPType = "COMMON_STAR"
	ECPTypeAgent      ECPType = "AGENT"
)

func (e ECPType) String() string {
	switch e {
	case ECPTypeShop:
		return "商家"
	case ECPTypeShopStar:
		return "商家达人"
	case ECPTypeCommonStar:
		return "普通达人"
	case ECPTypeAgent:
		return "百应机构"
	default:
		return "未知"
	}
}

type ExcludeLimitedRegion int

const (
	ExcludeLimitedRegionNo  ExcludeLimitedRegion = 0
	ExcludeLimitedRegionYes ExcludeLimitedRegion = 1
)

type DistrictType bool

const (
	DistrictTypeExclude DistrictType = true  //排除地域
	DistrictTypeInclude DistrictType = false //定向地域
)

type RetargetingTagsType int

const (
	RetargetingTagsTypeAll    RetargetingTagsType = 0
	RetargetingTagsTypeCustom RetargetingTagsType = 1
)

type OrderPlatform string

const (
	OrderPlatformAll       OrderPlatform = "ALL"       // 全部
	OrderPlatformQianchuan OrderPlatform = "QIANCHUAN" // 千川pc（默认）
	OrderPlatformEcpAweme  OrderPlatform = "ECP_AWEME" // 小店随心推
)

//time_granularity
//
//string
//
//时间粒度 ，如果不传，返回查询日期内的聚合数据
//允许值:
//TIME_GRANULARITY_DAILY (按天维度),会返回每天的数据
//TIME_GRANULARITY_HOURLY (按小时维度)，会返回每小时维度的数据

type TimeGranularity string

const (
	TimeGranularityDaily TimeGranularity = "TIME_GRANULARITY_DAILY"  // 按天维度
	TimeGranularityHour  TimeGranularity = "TIME_GRANULARITY_HOURLY" // 按小时维度
)

type OrderType string

const (
	OrderTypeDesc OrderType = "DESC" // 降序
	OrderTypeAsc  OrderType = "ASC"  // 升序
)

type AdGroupStatus string

const (
	AdGroupStatusAll     AdGroupStatus = "ALL"    // 所有包含已删除
	AdGroupStatusEnable  AdGroupStatus = "ENABLE" // 启用
	AdGroupStatusDisable AdGroupStatus = "DISABLE"
	AdGroupStatusDelete  AdGroupStatus = "DELETE"
)

func (a AdGroupStatus) String() string {
	switch a {
	case AdGroupStatusAll:
		return "所有包含已删除"
	case AdGroupStatusEnable:
		return "启用"
	case AdGroupStatusDisable:
		return "暂停"
	case AdGroupStatusDelete:
		return "已删除"
	default:
		return "未知状态"
	}
}
