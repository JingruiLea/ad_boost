package tools

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetCategoryList(ctx context.Context, adID int64) error {
	params := map[string]interface{}{
		"advertiser_id": 1748031128935424,
	}
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().Get(ctx, "https://ad.oceanengine.com/open_api/2/tools/aweme_multi_level_category/get/", httpclient.CommonHeader, &resp, params)
	if err != nil {
		logs.CtxErrorf(ctx, "GetCategoryList httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetCategoryList respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}
