package live_report

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

// https://ad.oceanengine.com/open_api/2/report/live_room/flow_category/get/
func GetFlowCategory(ctx context.Context, req *GetFlowCategoryReq) (*GetFlowCategoryRespData, error) {
	var resp GetFlowCategoryRespData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://ad.oceanengine.com/open_api/2/report/live_room/flow_category/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetFlowCategory httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	fmt.Printf("GetFlowCategory respMap: %s", utils.GetJsonStr(resp))
	return &resp, nil
}

type GetFlowCategoryReq struct {
	AdvertiserID int64     `json:"advertiser_id"`         // 广告主id
	StartTime    *string   `json:"start_time,omitempty"`  // 报表开始时间
	EndTime      *string   `json:"end_time,omitempty"`    // 报表结束时间
	Fields       []string  `json:"fields,omitempty"`      // 指标字段
	Filtering    Filtering `json:"filtering,omitempty"`   // 筛选字段
	OrderField   *string   `json:"order_field,omitempty"` // 排序指标
	OrderType    *string   `json:"order_type,omitempty"`  // 排序方式
}

// Filtering 筛选字段的结构体
type Filtering struct {
	ProductIDs []int `json:"product_ids,omitempty"` // 商品id列表
	RoomIDs    []int `json:"room_ids,omitempty"`    // 直播间id列表
}

type GetFlowCategoryRespData struct {
	Fields GetFlowCategoryFieldItem `json:"fields"` // 指标字段
}

type GetFlowCategoryFieldItem struct {
	LiveAvgWatchDuration float64 `json:"live_avg_watch_duration"` // 人均停留时长
	LiveFollowCount      int     `json:"live_follow_count"`       // 关注数
	LiveGiftCount        int     `json:"live_gift_count"`         // 打赏次数
	LiveGiftMoney        int     `json:"live_gift_money"`         // 礼物总金额，单位：音浪
	LiveWatchCount       int     `json:"live_watch_count"`        // 观看数
	LiveWatchUCount      int     `json:"live_watch_ucount"`       // 观看人数
	PayOrderCount        int     `json:"pay_order_count"`         // 商品订单数
	PayOrderGMV          float64 `json:"pay_order_gmv"`           // 商品订单金额
	FirstFlowCategory    string  `json:"first_flow_category"`     // 一级流量来源
}

type GetFlowCategoryField string

// 定义所有字段的常量
const (
	GetFlowCategoryFieldLiveAvgWatchDuration GetFlowCategoryField = "live_avg_watch_duration"
	GetFlowCategoryFieldLiveFollowCount      GetFlowCategoryField = "live_follow_count"
	GetFlowCategoryFieldLiveGiftCount        GetFlowCategoryField = "live_gift_count"
	GetFlowCategoryFieldLiveGiftMoney        GetFlowCategoryField = "live_gift_money"
	GetFlowCategoryFieldLiveWatchCount       GetFlowCategoryField = "live_watch_count"
	GetFlowCategoryFieldLiveWatchUCount      GetFlowCategoryField = "live_watch_ucount"
	GetFlowCategoryFieldPayOrderCount        GetFlowCategoryField = "pay_order_count"
	GetFlowCategoryFieldPayOrderGMV          GetFlowCategoryField = "pay_order_gmv"
	GetFlowCategoryFieldFirstFlowCategory    GetFlowCategoryField = "first_flow_category"
)

func (f GetFlowCategoryField) All() []GetFlowCategoryField {
	return []GetFlowCategoryField{
		GetFlowCategoryFieldLiveAvgWatchDuration,
		GetFlowCategoryFieldLiveFollowCount,
		GetFlowCategoryFieldLiveGiftCount,
		GetFlowCategoryFieldLiveGiftMoney,
		GetFlowCategoryFieldLiveWatchCount,
		GetFlowCategoryFieldLiveWatchUCount,
		GetFlowCategoryFieldPayOrderCount,
		GetFlowCategoryFieldPayOrderGMV,
		GetFlowCategoryFieldFirstFlowCategory,
	}
}
