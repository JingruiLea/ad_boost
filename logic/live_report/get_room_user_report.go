package live_report

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

// 获取直播间用户洞察
func GetRoomUserReport(ctx context.Context, req *GetRoomUserReportReq) error {
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://api.oceanengine.com/open_api/v1.0/qianchuan/today_live/room/user/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetReport httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetReport respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}

type GetRoomUserReportReq struct {
	AdvertiserID int64           `json:"advertiser_id"`         // 广告主id
	RoomID       int64           `json:"room_id"`               // 直播间ID
	ActionEvent  ActionEvent     `json:"action_event"`          //用户来源
	Dimension    []RoomDimension `json:"dimension"`             // 维度
	FlowSource   RoomFlowSource  `json:"flow_source,omitempty"` //广告类型
}

type ActionEvent string

const (
	//进入直播间
	ActionEventEnter = "ENTER"
	//支付成功
	ActionEventPay = "PAY"
)

type RoomDimension string

const (
	RoomDimensionCity   = "CITY"
	RoomDimensionGender = "GENDER"
	RoomDimensionAge    = "AGE"
)

type RoomFlowSource string

const (
	RoomFlowSourceAll = "ALL"
	RoomFlowSourcePc  = "PC"
)
