package talent

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetTalentByID(ctx context.Context, adID int64) error {
	params := map[string]interface{}{
		"advertiser_id": 1748031128935424,
		"label_ids":     []int{104651203101},
	}
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().Get(ctx, "https://ad.oceanengine.com/open_api/2/tools/aweme_author_info/get/", httpclient.CommonHeader, &resp, params)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccount httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetAdAccount respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}
