package account

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetPublicInfo(ctx context.Context, accountID int64) error {
	params := map[string]interface{}{
		"advertiser_id": accountID,
	}
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().AdGet(ctx, accountID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/finance/wallet/get/", &resp, params)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccount httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetAdAccount respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}
