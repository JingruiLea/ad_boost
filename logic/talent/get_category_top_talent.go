package talent

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetCategoryTopTalent(ctx context.Context, adID int64, categoryID int64) error {
	var params = map[string]interface{}{
		"advertiser_id": 1748031128935424,
		"category_id":   categoryID,
		"behaviors":     ttypes.AwemeFanBehaviorsCommentedUser.Common(),
	}
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().Get(ctx, "https://ad.oceanengine.com/open_api/2/tools/aweme_category_top_author/get/", &resp, params)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccount httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetAdAccount respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}
