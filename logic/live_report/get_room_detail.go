package live_report

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetRoomDetail(ctx context.Context, req *GetRoomDetailReq) error {
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://api.oceanengine.com/open_api/v1.0/qianchuan/today_live/room/detail/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetReport httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetReport respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}

type GetRoomDetailReq struct {
	AdvertiserID int64 `json:"advertiser_id"` // 广告主id
	RoomID       int64 `json:"room_id"`       // 直播间ID
}
