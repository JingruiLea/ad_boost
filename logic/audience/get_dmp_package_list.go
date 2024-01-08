package audience

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

type GetDMPPackageListReq struct {
	AdvertiserId        int64                      `json:"advertiser_id"`
	RetargetingTagsType ttypes.RetargetingTagsType `json:"retargeting_tags_type"`
	Offset              int                        `json:"offset,omitempty"`
	Limit               int                        `json:"limit,omitempty"`
}

func GetDMPPackageList(ctx context.Context, req *GetDMPPackageListReq) error {
	var resp GetDMPPackageListRespData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserId, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/dmp/audiences/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetDMPPackageList httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetDMPPackageList respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}

type GetDMPPackageListRespData struct {
	Offset          int               `json:"offset"`
	RetargetingTags []*RetargetingTag `json:"retargeting_tags"`
	TotalNum        int               `json:"total_num"`
}

type RetargetingTag struct {
	CoverNum           int         `json:"cover_num"`
	HasOfflineTag      int         `json:"has_offline_tag"`
	IsCommon           int         `json:"is_common"`
	Name               string      `json:"name"`
	RetargetingTagsId  int         `json:"retargeting_tags_id"`
	RetargetingTagsOp  string      `json:"retargeting_tags_op"`
	RetargetingTagsTip interface{} `json:"retargeting_tags_tip"`
	Source             string      `json:"source"`
	Status             int         `json:"status"`
}
