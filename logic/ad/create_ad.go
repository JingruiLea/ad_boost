package ad

import (
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/bo"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
	"golang.org/x/net/context"
)

func CreateAd(ctx context.Context, ad *bo.CreateAd) (adID int64, err error) {
	url := "https://api.oceanengine.com/open_api/v1.0/qianchuan/ad/create/"
	var resp CreateAdRespData
	fmt.Printf("AdCreate reqMap: %s", utils.GetJsonStr(ad))
	err = httpclient.NewClient().AdPost(ctx, ad.AdvertiserID, url, ad, &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "AdCreate httpclient.NewClient().Post error: %v", err)
		return 0, err
	}
	return resp.AdId, nil
}

type CreateAdRespData struct {
	AdId        int64         `json:"ad_id"`
	NoticeInfos []interface{} `json:"notice_infos"`
}
