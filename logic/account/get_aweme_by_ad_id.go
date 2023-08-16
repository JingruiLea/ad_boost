package account

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
	"gorm.io/datatypes"
)

func GetAwemeByAdID(ctx context.Context, adID int64, page, size int32) (*GetAwemeByAdIDRespData, error) {
	var resp GetAwemeByAdIDResp
	params := map[string]interface{}{
		"advertiser_id": adID,
		"page":          page,
		"page_size":     size,
	}
	err := httpclient.NewClient().Get(ctx, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/aweme/authorized/get/", httpclient.CommonHeader, &resp, params)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccountByShopID httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	fmt.Printf("GetAdAccountByShopID respMap: %s", utils.GetJsonStr(resp))
	return resp.Data, err
}

type GetAwemeByAdIDResp struct {
	ttypes.BaseResp
	Data *GetAwemeByAdIDRespData `json:"data"`
}

type GetAwemeByAdIDRespData struct {
	AwemeIdList []*AwemeInfo     `json:"aweme_id_list"`
	PageInfo    *ttypes.PageInfo `json:"page_info"`
}

type AwemeInfo struct {
	AwemeAvatar             string   `json:"aweme_avatar"`
	AwemeHasLivePermission  bool     `json:"aweme_has_live_permission"`
	AwemeHasUniProm         bool     `json:"aweme_has_uni_prom"`
	AwemeHasVideoPermission bool     `json:"aweme_has_video_permission"`
	AwemeId                 int64    `json:"aweme_id"`
	AwemeName               string   `json:"aweme_name"`
	AwemeShowId             string   `json:"aweme_show_id"`
	AwemeStatus             string   `json:"aweme_status"`
	BindType                []string `json:"bind_type"`
}

func (a *AwemeInfo) ToModel() *model.Aweme {
	ret := &model.Aweme{
		AwemeAvatar:             a.AwemeAvatar,
		AwemeHasLivePermission:  a.AwemeHasLivePermission,
		AwemeHasUniProm:         a.AwemeHasUniProm,
		AwemeHasVideoPermission: a.AwemeHasVideoPermission,
		AwemeId:                 a.AwemeId,
		AwemeName:               a.AwemeName,
		AwemeShowId:             a.AwemeShowId,
		AwemeStatus:             a.AwemeStatus,
		BindType:                datatypes.JSON(utils.GetJsonStr(a.BindType)),
	}
	return ret
}
