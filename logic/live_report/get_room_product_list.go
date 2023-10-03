package live_report

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetRoomProductList(ctx context.Context, req *GetRoomProductListReq) error {
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://api.oceanengine.com/open_api/v1.0/qianchuan/today_live/room/product_list/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetReport httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetReport respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}

type GetRoomProductListReq struct {
	AdvertiserID  int64              `json:"advertiser_id"`  // 广告主id
	RoomID        int64              `json:"room_id"`        // 直播间ID
	Fields        []RoomMetricsField `json:"fields"`         // 维度
	ExplainStatus ExplainStatus      `json:"explain_status"` // 是否展示商品状态
	Page          int                `json:"page"`           // 页码
	PageSize      int                `json:"page_size"`      // 每页数量
}

type ExplainStatus string

const (
	ExplainStatusAll        = "ALL"
	ExplainStatusHasExplain = "HASEXPLAIN"
	ExplainStatusUnExplain  = "UNEXPLAIN"
)
