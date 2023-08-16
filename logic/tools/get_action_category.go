package tools

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetActionCategory(ctx context.Context, adID int64) error {
	params := map[string]interface{}{
		"advertiser_id": 1748031128935424,
		"action_scene":  []string{"E-COMMERCE"},
		"action_days":   180,
	}
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().Get(ctx, "https://ad.oceanengine.com/open_api/2/tools/interest_action/action/category/", httpclient.CommonHeader, &resp, params)
	if err != nil {
		logs.CtxErrorf(ctx, "GetCategoryList httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetCategoryList respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}
