package ad

import (
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
	"golang.org/x/net/context"
)

func CreateAd(ctx context.Context, ad *ttypes.Ad) (adID int64, err error) {
	url := "https://api.oceanengine.com/open_api/v1.0/qianchuan/ad/create/"
	var resp CreateAdResp
	fmt.Printf("AdCreate reqMap: %s", utils.GetJsonStr(ad))
	err = httpclient.NewClient().Post(ctx, url, httpclient.CommonHeader, ad, &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "AdCreate httpclient.NewClient().Post error: %v", err)
		return 0, err
	}
	if resp.Code != 0 || resp.Data == nil {
		logs.CtxErrorf(ctx, "AdCreate resp.Code != 0, resp: %s", utils.GetJsonStr(resp))
		return 0, fmt.Errorf("AdCreate resp.Code != 0, resp: %s", utils.GetJsonStr(resp))
	}
	return resp.Data.AdId, nil
}

type CreateAdResp struct {
	ttypes.BaseResp
	Data *CreateAdRespData `json:"data,omitempty"`
}

type CreateAdRespData struct {
	AdId        int64         `json:"ad_id"`
	NoticeInfos []interface{} `json:"notice_infos"`
}
