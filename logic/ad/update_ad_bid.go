package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func UpdateAdBid(ctx context.Context, advertiserID, adID int64) error {
	//   "list": [
	//      1748031128935424,
	//      1748031128935424,
	//      1767935594672136,
	//      1769126587284494
	//    ],
	var req UpdateAdBidReq
	req.AdvertiserId = 1748031128935424

	var resp = make(map[string]interface{})
	err := httpclient.NewClient().Post(ctx, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/ad/bid/update/", httpclient.CommonHeader, req, &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "UpdateAdBid httpclient.NewClient().Post error: %v", err)
		return err
	}
	fmt.Printf("UpdateAdBid respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}

type UpdateAdBidReq struct {
	AdvertiserId int64  `json:"advertiser_id"`
	Data         []*Bid `json:"data"`
}

type Bid struct {
	AdId int64   `json:"ad_id"`
	Bid  float32 `json:"bid"`
}
