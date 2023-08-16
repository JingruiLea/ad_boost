package live_report

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetTodayLiveRoom(ctx context.Context, req *GetTodayLiveRoomReq) error {
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().Get(ctx, "https://api.oceanengine.com/open_api/v1.0/qianchuan/today_live/room/get/", httpclient.CommonHeader, &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetReport httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetReport respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}

type GetTodayLiveRoomReq struct {
	AdvertiserID int64              `json:"advertiser_id"` // 广告主id
	AwemeID      int64              `json:"aweme_id"`      // 抖音号ID
	DateTime     string             `json:"date_time"`     // 日期
	RoomStatus   RoomStatus         `json:"room_status"`   // 直播间状态
	AdStatus     AdStatus           `json:"ad_status"`     // 广告状态
	Fields       []RoomMetricsField `json:"fields"`        // 需要查询的消耗指标
	Page         int                `json:"page"`          // 页码
	PageSize     int                `json:"page_size"`     // 每页数量
}

type RoomStatus string

const (
	RoomStatusAll    = "ALL"
	RoomStatusLiving = "LIVING"
	RoomStatusFinish = "FINISH"
)

type AdStatus string

const (
	AdStatusAll        = "ALL"
	AdStatusDeliveryOk = "DELIVERY_OK"
	AdStatusNoDelivery = "NO_DELIVERY"
)
