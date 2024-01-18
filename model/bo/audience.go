package bo

import (
	"github.com/JingruiLea/ad_boost/model/ttypes"
)

// Audience 定向人群设置
type Audience struct {
	OrientationName         string                       `json:"orientation_name,omitempty"` // 定向包名称
	AudienceMode            ttypes.AudienceMode          `json:"audience_mode"`              // 人群定向模式
	OrientationID           int64                        `json:"orientation_id"`             // 定向包id；仅专业推广支持，极速推广不支持
	ExcludeLimitedRegion    ttypes.ExcludeLimitedRegion  `json:"exclude_limited_region"`     // 排除限运地区，0：否，默认值；1：是
	DistrictType            ttypes.DistrictType          `json:"district_type"`              // 定向or排除地域；true：排除地域；false：定向地域；仅通投广告+极速推广+托管计划，支持
	District                ttypes.District              `json:"district"`                   // 地域定向类型，配合city字段使用；枚举值：CITY：省市，COUNTY：区县，NONE：不限；默认值：NONE
	City                    []int                        `json:"city"`                       // 具体定向的城市列表
	LocationType            ttypes.LocationType          `json:"location_type"`              // 地域定向的用户状态类型；枚举值：CURRENT，HOME，TRAVEL，ALL
	Gender                  ttypes.Gender                `json:"gender"`                     // 性别；枚举值: GENDER_FEMALE，GENDER_MALE，NONE
	Age                     []*ttypes.Age                `json:"age"`                        // 年龄区间；枚举值：AGE_BETWEEN_18_23, AGE_BETWEEN_24_30, AGE_BETWEEN_31_40, AGE_BETWEEN_41_49, AGE_ABOVE_50
	AwemeFanBehaviors       []ttypes.AwemeFanBehavior    `json:"aweme_fan_behaviors"`        // 抖音用户行为类型
	AwemeFanBehaviorsDays   ttypes.AwemeFanBehaviorsDays `json:"aweme_fan_behaviors_days"`   // 抖音达人互动用户行为天数；枚举值：DAYS_15，DAYS_30，DAYS_60
	AwemeFanCategories      []int                        `json:"aweme_fan_categories"`       // 抖音达人分类ID列表
	AwemeFanAccounts        []int                        `json:"aweme_fan_accounts"`         // 抖音达人ID列表
	AutoExtendEnabled       int                          `json:"auto_extend_enabled"`        // 是否启用智能放量
	AutoExtendTargets       []string                     `json:"auto_extend_targets"`        // 可放开定向列表
	Platform                []string                     `json:"platform"`                   // 投放平台列表
	SmartInterestAction     string                       `json:"smart_interest_action"`      // 行为兴趣意向定向模式
	ActionScene             []string                     `json:"action_scene"`               // 行为场景
	ActionDays              int                          `json:"action_days"`                // 用户发生行为天数
	ActionCategories        []int                        `json:"action_categories"`          // 行为类目词
	ActionWords             []int                        `json:"action_words"`               // 行为关键词
	InterestCategories      []int                        `json:"interest_categories"`        // 兴趣类目词
	InterestWords           []int                        `json:"interest_words"`             // 兴趣关键词
	AC                      []string                     `json:"ac"`                         // 网络类型
	RetargetingTagsInclude  []int                        `json:"retargeting_tags_include"`   // 定向人群包id列表
	RetargetingTagsExclude  []int                        `json:"retargeting_tags_exclude"`   // 排除人群包id列表
	LivePlatformTags        []string                     `json:"live_platform_tags"`         // 平台精选人群包
	InactiveRetargetingTags []InactiveTag                `json:"inactive_retargeting_tags"`  // 失效的人群包列表
	NewCustomer             string                       `json:"new_customer"`               // 新客定向；枚举值：NO_BUY，NONE，NO_BUY_BRAND，NO_BUY_DOUYIN
}

// InactiveTag 描述失效的人群包
type InactiveTag struct {
	RetargetingTag int64  `json:"retargeting_tag"` // 人群包id
	Name           string `json:"name"`            // 人群包名称
	InactiveType   string `json:"inactive_type"`   // 失效类型；枚举值：EXPIRE，TAG_OFFLINE，MANUAL_OFFLINE
}

func NewAudience() *Audience {
	return &Audience{
		OrientationID:          0,
		AudienceMode:           ttypes.AudienceModeCustom,
		ExcludeLimitedRegion:   1,
		DistrictType:           false,
		District:               "",
		City:                   nil,
		LocationType:           "",
		Gender:                 "",
		Age:                    nil,
		AwemeFanBehaviors:      ttypes.AwemeFanBehaviorsCommentedUser.Common(),
		AwemeFanBehaviorsDays:  ttypes.AwemeFanBehaviorsDays60,
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
		RetargetingTagsInclude: nil,
		RetargetingTagsExclude: nil,
		LivePlatformTags:       nil,
		NewCustomer:            "",
	}
}

func (a *Audience) WithOrientationID(orientationId int64) *Audience {
	a.OrientationID = orientationId
	return a
}
