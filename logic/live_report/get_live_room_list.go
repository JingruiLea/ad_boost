package live_report

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
	"time"
)

func GetNowLiveRoomList(ctx context.Context, accountID, awemeID int64) (*GetLiveRoomListRespData, error) {
	var resp GetLiveRoomListRespData
	var req = &GetLiveRoomListReq{
		AdvertiserID: accountID,
		AwemeID:      awemeID,
		DateTime:     time.Now().Format("2006-01-02"),
		RoomStatus:   RoomStatusLiving,
		AdStatus:     AdStatusAll,
		Fields:       RoomMetricsFieldStatCost.All(),
		Page:         1,
		PageSize:     10,
	}
	err := httpclient.NewClient().AdGet(ctx, accountID, "https://api.oceanengine.com/open_api/v1.0/qianchuan/today_live/room/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetReport httpclient.NewClient().AdGet error: %v", err)
		return nil, err
	}
	fmt.Printf("GetReport respMap: \n\n%s\n\n", utils.GetJsonStr(resp))
	return &resp, err
}

func GetLiveRoomList(ctx context.Context, req *GetLiveRoomListReq) (*GetLiveRoomListRespData, error) {
	var resp GetLiveRoomListRespData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://api.oceanengine.com/open_api/v1.0/qianchuan/today_live/room/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetReport httpclient.NewClient().AdGet error: %v", err)
		return nil, err
	}
	fmt.Printf("GetReport respMap: \n\n%s\n\n", utils.GetJsonStr(resp))
	return &resp, err
}

type GetLiveRoomListReq struct {
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

type GetLiveRoomListRespData struct {
	PageInfo ttypes.PageInfo `json:"page_info"`
	RoomList []*RoomReport   `json:"room_list"`
}

type RoomReport struct {
	AdLiveCreateOrderRate         float64  `json:"ad_live_create_order_rate"`
	AdLivePayOrderGmvAvg          int      `json:"ad_live_pay_order_gmv_avg"`
	AdLivePayOrderRate            float64  `json:"ad_live_pay_order_rate"`
	AdLivePrepayAndPayOrderGmvRoi float64  `json:"ad_live_prepay_and_pay_order_gmv_roi"`
	AwemeAvatar                   []string `json:"aweme_avatar"`
	AwemeId                       int64    `json:"aweme_id"`
	AwemeName                     string   `json:"aweme_name"`
	ClickCnt                      int      `json:"click_cnt"`
	ConvertCnt                    int      `json:"convert_cnt"`
	ConvertRate                   float64  `json:"convert_rate"`
	CpaPlatform                   float64  `json:"cpa_platform"`
	CpcPlatform                   float64  `json:"cpc_platform"`
	CpmPlatform                   float64  `json:"cpm_platform"`
	Ctr                           float64  `json:"ctr"`
	EndTime                       string   `json:"end_time"`
	LiveClickCartCountAlias       int      `json:"live_click_cart_count_alias"`
	LiveClickProductCountAlias    int      `json:"live_click_product_count_alias"`
	LiveCreateOrderCountAlias     int      `json:"live_create_order_count_alias"`
	LiveCreateOrderRate           float64  `json:"live_create_order_rate"`
	LiveOrderPayCouponAmount      int      `json:"live_order_pay_coupon_amount"`
	LivePayOrderCountAlias        int      `json:"live_pay_order_count_alias"`
	LivePayOrderGmvAlias          float64  `json:"live_pay_order_gmv_alias"`
	LivePayOrderGmvAvg            float64  `json:"live_pay_order_gmv_avg"`
	LivePayOrderGmvRoi            float64  `json:"live_pay_order_gmv_roi"`
	LivePayOrderRate              float64  `json:"live_pay_order_rate"`
	LivePrepayOrderCountAlias     int      `json:"live_prepay_order_count_alias"`
	LivePrepayOrderGmvAlias       int      `json:"live_prepay_order_gmv_alias"`
	LiveWatchOneMinuteCount       int      `json:"live_watch_one_minute_count"`
	LubanLiveOrderCount           int      `json:"luban_live_order_count"`
	LubanLivePayOrderCount        int      `json:"luban_live_pay_order_count"`
	LubanLivePayOrderGmv          float64  `json:"luban_live_pay_order_gmv"`
	LubanLivePayOrderGpm          float64  `json:"luban_live_pay_order_gpm"`
	LubanLivePrepayOrderCount     int      `json:"luban_live_prepay_order_count"`
	LubanLivePrepayOrderGmv       float64  `json:"luban_live_prepay_order_gmv"`
	RoomCover                     []string `json:"room_cover"`
	RoomDelivery                  int      `json:"room_delivery"`
	RoomId                        int64    `json:"room_id"`
	RoomStatus                    string   `json:"room_status"`
	RoomTitle                     string   `json:"room_title"`
	StartTime                     string   `json:"start_time"`
	StatCost                      float64  `json:"stat_cost"`
	TotalLiveCommentCnt           int      `json:"total_live_comment_cnt"`
	TotalLiveFansClubJoinCnt      int      `json:"total_live_fans_club_join_cnt"`
	TotalLiveFollowCnt            int      `json:"total_live_follow_cnt"`
	TotalLiveGiftAmount           float64  `json:"total_live_gift_amount"`
	TotalLiveGiftCnt              int      `json:"total_live_gift_cnt"`
	TotalLivePayOrderGpm          float64  `json:"total_live_pay_order_gpm"`
	TotalLiveShareCnt             int      `json:"total_live_share_cnt"`
	TotalLiveWatchCnt             int      `json:"total_live_watch_cnt"`
}
