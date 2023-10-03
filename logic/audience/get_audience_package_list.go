package audience

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
	"gorm.io/datatypes"
)

func GetAudiencePackageList(ctx context.Context, req *GetAudiencePackageListReq) (*GetAudiencePackageListRespData, error) {
	var resp GetAudiencePackageListRespData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://api.oceanengine.com/open_api/v1.0/qianchuan/orientation_package/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetAudiencePackageList httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	fmt.Printf("GetAudiencePackageList respMap: %s", utils.GetJsonStr(resp))
	return &resp, nil
}

type GetAudiencePackageListReq struct {
	AdvertiserID int64                  `json:"advertiser_id"`       // 千川广告主账户ID
	Filtering    *AudiencePackageFilter `json:"filtering,omitempty"` // 过滤条件
	Page         int                    `json:"page"`                // 页码
	PageSize     int                    `json:"page_size"`           // 页面大小
}

type AudiencePackageFilter struct {
	Name string  `json:"name"` // 定向包名称
	IDs  []int64 `json:"id"`   // 定向包ID
}

type GetAudiencePackageListResp struct {
	ttypes.BaseResp
	Data *GetAudiencePackageListRespData `json:"data"`
}

type GetAudiencePackageListRespData struct {
	List     []*AudienceConfig `json:"list"`
	PageInfo *ttypes.PageInfo  `json:"page_info"`
}

type AudienceConfig struct {
	InActiveRetargetingTags interface{} `json:"InActive_retargeting_tags"`
	Ac                      []string    `json:"ac"`
	ActionCategories        []int       `json:"action_categories"`
	ActionDays              int         `json:"action_days"`
	ActionScene             []string    `json:"action_scene"`
	ActionWords             []int       `json:"action_words"`
	AdNum                   int         `json:"ad_num"`
	Age                     []string    `json:"age"`
	AutoExtendEnabled       int         `json:"auto_extend_enabled"`
	AutoExtendTargets       []string    `json:"auto_extend_targets"`
	AwemeFanAccounts        []int64     `json:"aweme_fan_accounts"`
	AwemeFanBehaviors       []string    `json:"aweme_fan_behaviors"`
	AwemeFanBehaviorsDays   string      `json:"aweme_fan_behaviors_days"`
	AwemeFanCategories      []int       `json:"aweme_fan_categories"`
	City                    []int       `json:"city"`
	District                *string     `json:"district"`
	Gender                  *string     `json:"gender"`
	InterestCategories      []int       `json:"interest_categories"`
	InterestWords           interface{} `json:"interest_words"`
	LivePlatformTags        []string    `json:"live_platform_tags"`
	LocationType            *string     `json:"location_type"`
	NewCustomer             string      `json:"new_customer"`
	OrientationID           int64       `json:"orientation_id"`
	OrientationInfo         interface{} `json:"orientation_info"`
	OrientationName         string      `json:"orientation_name"`
	Platform                interface{} `json:"platform"`
	RetargetingTagsExclude  []int       `json:"retargeting_tags_exclude"`
	RetargetingTagsInclude  []int       `json:"retargeting_tags_include"`
	SmartInterestAction     string      `json:"smart_interest_action"`
}

func (a *AudienceConfig) ToModel() *model.Audience {
	ret := &model.Audience{
		AudienceID: a.OrientationID,
		Name:       a.OrientationName,
		Config:     datatypes.JSON(utils.GetJsonStr(a)),
	}
	return ret
}
