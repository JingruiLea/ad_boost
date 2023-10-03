package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetAdDetail(ctx context.Context, advertiserID, adID int64) error {
	var req = map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
	}

	var resp = make(map[string]interface{})
	err := httpclient.NewClient().AdGet(ctx, advertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/ad/detail/get/", &resp, req)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccount httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetAdAccount respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}

type T struct {
	Code int `json:"code"`
	Data struct {
		AdCreateTime string `json:"ad_create_time"`
		AdId         int64  `json:"ad_id"`
		AdModifyTime string `json:"ad_modify_time"`
		Audience     struct {
			ActionCategories        []int         `json:"action_categories"`
			ActionDays              int           `json:"action_days"`
			ActionScene             []string      `json:"action_scene"`
			Age                     []string      `json:"age"`
			AudienceMode            string        `json:"audience_mode"`
			AutoExtendEnabled       int           `json:"auto_extend_enabled"`
			District                string        `json:"district"`
			ExcludeLimitedRegion    int           `json:"exclude_limited_region"`
			Gender                  string        `json:"gender"`
			InactiveRetargetingTags []interface{} `json:"inactive_retargeting_tags"`
			InterestCategories      []int         `json:"interest_categories"`
			LivePlatformTags        []interface{} `json:"live_platform_tags"`
			NewCustomer             string        `json:"new_customer"`
			SmartInterestAction     string        `json:"smart_interest_action"`
		} `json:"audience"`
		AwemeInfo []struct {
			AwemeAvatar string `json:"aweme_avatar"`
			AwemeId     int64  `json:"aweme_id"`
			AwemeName   string `json:"aweme_name"`
			AwemeShowId string `json:"aweme_show_id"`
		} `json:"aweme_info"`
		CampaignId           int64  `json:"campaign_id"`
		CampaignScene        string `json:"campaign_scene"`
		CreativeAutoGenerate int    `json:"creative_auto_generate"`
		CreativeList         []struct {
			CreativeCreateTime string `json:"creative_create_time"`
			CreativeId         int64  `json:"creative_id"`
			CreativeModifyTime string `json:"creative_modify_time"`
			ImageMode          string `json:"image_mode"`
		} `json:"creative_list"`
		CreativeMaterialMode string `json:"creative_material_mode"`
		DeliverySetting      struct {
			AllowQcpx          bool   `json:"allow_qcpx"`
			Budget             int    `json:"budget"`
			BudgetMode         string `json:"budget_mode"`
			DeepBidType        string `json:"deep_bid_type"`
			DeepExternalAction string `json:"deep_external_action"`
			EndTime            string `json:"end_time"`
			ExternalAction     string `json:"external_action"`
			LiveScheduleType   string `json:"live_schedule_type"`
			ProductNewOpen     bool   `json:"product_new_open"`
			RoiGoal            int    `json:"roi_goal"`
			ScheduleFixedRange int    `json:"schedule_fixed_range"`
			ScheduleTime       string `json:"schedule_time"`
			SmartBidType       string `json:"smart_bid_type"`
			StartTime          string `json:"start_time"`
			VideoScheduleType  string `json:"video_schedule_type"`
		} `json:"delivery_setting"`
		DynamicCreative int           `json:"dynamic_creative"`
		FirstIndustryId int           `json:"first_industry_id"`
		IsHomepageHide  int           `json:"is_homepage_hide"`
		Keywords        []interface{} `json:"keywords"`
		LabAdType       string        `json:"lab_ad_type"`
		MarketingGoal   string        `json:"marketing_goal"`
		MarketingScene  string        `json:"marketing_scene"`
		Name            string        `json:"name"`
		OptStatus       string        `json:"opt_status"`
		PivativeWords   struct {
		} `json:"pivative_words"`
		ProductInfo              []interface{} `json:"product_info"`
		ProgrammaticCreativeCard struct {
		} `json:"programmatic_creative_card"`
		ProgrammaticCreativeTitleList []interface{} `json:"programmatic_creative_title_list"`
		RoomInfo                      []struct {
			AnchorAvatar string `json:"anchor_avatar"`
			AnchorId     int64  `json:"anchor_id"`
			AnchorName   string `json:"anchor_name"`
			RoomStatus   string `json:"room_status"`
		} `json:"room_info"`
		SecondIndustryId int    `json:"second_industry_id"`
		Status           string `json:"status"`
		ThirdIndustryId  int    `json:"third_industry_id"`
		TrackUrl         struct {
		} `json:"track_url"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}
