package account

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/shop_dal"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetAdAccountByShopID(ctx context.Context, shopID int64) (data []int64, err error) {
	var resp GetAdAccountByShopIDResp
	shop, err := shop_dal.GetShopByShopID(ctx, shopID)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccountByShopID shop_dal.GetShopByShopID error: %v", err)
		return nil, err
	}
	err = httpclient.NewClient().Get(ctx, fmt.Sprintf("https://ad.oceanengine.com/open_api/v1.0/qianchuan/shop/advertiser/list/?shop_id=%d", shopID), map[string]string{
		"Access-Token": shop.AccessToken,
		"Content-Type": "application/json",
	}, &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccountByShopID httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	fmt.Printf("GetAdAccountByShopID respMap: %s", utils.GetJsonStr(resp))
	if resp.Code != 0 || resp.Data.List == nil {
		logs.CtxErrorf(ctx, "GetAdAccountByShopID resp.Code != 0, resp: %s", utils.GetJsonStr(resp))
		return nil, fmt.Errorf("GetAdAccountByShopID resp.Code != 0, resp: %s", utils.GetJsonStr(resp))
	}
	return resp.Data.List, nil
}

type GetAdAccountByShopIDResp struct {
	ttypes.BaseResp
	Data GetAdAccountByShopIDRespData `json:"data"`
}

type GetAdAccountByShopIDRespData struct {
	AdvIdList []struct {
		AdvId           int64         `json:"adv_id"`
		ExtraPermission []interface{} `json:"extra_permission"`
	} `json:"adv_id_list"`
	List     []int64          `json:"list"`
	PageInfo *ttypes.PageInfo `json:"page_info"`
}
