package account

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetAccountType(ctx context.Context, accountIDs []int64) (*GetAccountTypeRespData, error) {
	req := map[string]interface{}{
		"advertiser_ids": accountIDs,
	}
	var resp GetAccountTypeRespData
	err := httpclient.NewClient().AdGet(ctx, accountIDs[0], "https://api.oceanengine.com/open_api/v1.0/qianchuan/advertiser/type/get/", &resp, req)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAccountType httpclient.NewClient().AdGet error: %v", err)
		return nil, err
	}
	return &resp, err
}

type GetAccountTypeRespData struct {
	List []*GetAccountTypeRespDataItem `json:"list"`
}

type GetAccountTypeRespDataItem struct {
	ECPType      ttypes.ECPType `json:"ecp_type"`
	AdvertiserID int64          `json:"advertiser_id"`
}
