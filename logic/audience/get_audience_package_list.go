package audience

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/bo"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetAudiencePackageList(ctx context.Context, req *GetAudiencePackageListReq) (*GetAudiencePackageListRespData, error) {
	var resp GetAudiencePackageListRespData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://api.oceanengine.com/open_api/v1.0/qianchuan/orientation_package/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetAudiencePackageList httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	fmt.Printf("GetAudiencePackageList respMap: %s", utils.GetJsonStr(resp))
	return &resp, nil
}

type GetAudiencePackageListReq struct {
	AdvertiserID int64                  `json:"advertiser_id"`       // 千川广告主账户ID
	Filtering    *AudiencePackageFilter `json:"filtering,omitempty"` // 过滤条件
	Page         int                    `json:"page"`                // 页码
	PageSize     int                    `json:"page_size"`           // 页面大小
}

type AudiencePackageFilter struct {
	Name string  `json:"name"` // 定向包名称
	IDs  []int64 `json:"id"`   // 定向包ID
}

type GetAudiencePackageListResp struct {
	ttypes.BaseResp
	Data *GetAudiencePackageListRespData `json:"data"`
}

type GetAudiencePackageListRespData struct {
	List     []*bo.Audience   `json:"list"`
	PageInfo *ttypes.PageInfo `json:"page_info"`
}
