package account

import (
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
	"golang.org/x/net/context"
)

func GetShopAccount(ctx context.Context, at string) (accounts []*Account, err error) {
	var resp GetAccountResp
	err = httpclient.NewClient().Get(ctx, fmt.Sprintf("https://ad.oceanengine.com/open_api/oauth2/advertiser/get/?access_token=%s", at), httpclient.CommonHeader, &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "GetShopAccount httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	logs.CtxDebugf(ctx, "GetShopAccount respMap: %v", utils.GetJsonStr(resp))
	if resp.Code != 0 || resp.Data.List == nil {
		logs.CtxErrorf(ctx, "GetShopAccount resp.Code != 0, resp: %s", utils.GetJsonStr(resp))
		return nil, fmt.Errorf("GetShopAccount resp.Code != 0, resp: %s", utils.GetJsonStr(resp))
	}
	return resp.Data.List, nil
}

type GetAccountResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []*Account `json:"list"`
	} `json:"data"`
}

type Account struct {
	AdvertiserId   int    `json:"advertiser_id"` //这里其实是店铺ID
	AdvertiserName string `json:"advertiser_name"`
	IsValid        bool   `json:"is_valid"`
	AccountRole    string `json:"account_role"`
}
