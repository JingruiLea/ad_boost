package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/bo"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetAdDetail(ctx context.Context, advertiserID, adID int64) (*GetAdDetailRespData, error) {
	var req = map[string]interface{}{
		"advertiser_id": advertiserID,
		"ad_id":         adID,
	}

	var resp GetAdDetailRespData
	err := httpclient.NewClient().AdGet(ctx, advertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/ad/detail/get/", &resp, req)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccount httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	fmt.Printf("GetAdAccount respMap: %s", utils.GetJsonStr(resp))
	return &resp, nil
}

type GetAdDetailRespData struct {
	AdCreateTime         string                `json:"ad_create_time"`
	AdId                 int64                 `json:"ad_id"`
	AdModifyTime         string                `json:"ad_modify_time"`
	Audience             bo.Audience           `json:"audience"`
	AwemeInfo            []*bo.AwemeInfo       `json:"aweme_info"`
	CampaignId           int64                 `json:"campaign_id"`
	CampaignScene        string                `json:"campaign_scene"`
	CreativeAutoGenerate int                   `json:"creative_auto_generate"`
	CreativeList         []*bo.Creative        `json:"creative_list"`
	CreativeMaterialMode string                `json:"creative_material_mode"`
	DeliverySetting      *DeliverySetting      `json:"delivery_setting"`
	DynamicCreative      int                   `json:"dynamic_creative"`
	FirstIndustryId      int                   `json:"first_industry_id"`
	IsHomepageHide       int                   `json:"is_homepage_hide"`
	Keywords             []interface{}         `json:"keywords"`
	LabAdType            ttypes.LabAdType      `json:"lab_ad_type"`
	MarketingGoal        ttypes.MarketingGoal  `json:"marketing_goal"`
	MarketingScene       ttypes.MarketingScene `json:"marketing_scene"`
	Name                 string                `json:"name"`
	OptStatus            ttypes.OptStatus      `json:"opt_status"`
	PivativeWords        struct {
		PreciseWords []string `json:"precise_words"`
		PhraseWords  []string `json:"phrase_words"`
	} `json:"pivative_words"`
	ProductInfo                   []*bo.ProductInfo               `json:"product_info"`
	ProgrammaticCreativeCard      bo.ProgrammaticCreativeCard     `json:"programmatic_creative_card"`
	ProgrammaticCreativeTitleList []*bo.ProgrammaticCreativeTitle `json:"programmatic_creative_title_list"`
	RoomInfo                      []*bo.RoomInfo                  `json:"room_info"`
	SecondIndustryId              int                             `json:"second_industry_id"`
	Status                        ttypes.AdStatus                 `json:"status"`
	ThirdIndustryId               int                             `json:"third_industry_id"`
}
