package model

import (
	"github.com/JingruiLea/ad_boost/utils"
	"gorm.io/gorm"
	"reflect"
)

type AdReportItem struct {
	ID                   int64   `json:"id"`                     // 自增id
	AdID                 int64   `json:"ad_id" display:"广告计划ID"` // 广告计划id
	RoomID               int64   `json:"room_id" display:"直播间ID"`
	AdName               string  `json:"ad_name" display:"广告计划名称"` // 广告计划名称
	AdvertiserId         int64   `json:"advertiser_id" display:"广告主ID"`
	ClickCnt             int     `json:"click_cnt" display:"点击数"`   // 点击数
	ConvertCnt           int     `json:"convert_cnt" display:"转化数"` // 转化数
	ConvertCost          float64 `json:"convert_cost" display:"转化成本"`
	ConvertRate          float64 `json:"convert_rate" display:"转化率"`
	CpmPlatform          float64 `json:"cpm_platform" display:"平台千次展示费用"`
	Ctr                  float64 `json:"ctr" display:"点击率"`
	DyFollow             int     `json:"dy_follow" display:"新增粉丝数"`
	PayOrderAmount       float64 `json:"pay_order_amount" display:"直接成交金额"`          // 直接成交金额
	PayOrderCount        int     `json:"pay_order_count" display:"直接成交笔数"`           // 直接成交笔数
	PrepayAndPayOrderRoi float64 `json:"prepay_and_pay_order_roi" display:"直接支付roi"` // 直接支付roi
	ShowCnt              int     `json:"show_cnt" display:"展示次数"`                    // 展示次数
	StatCost             float64 `json:"stat_cost" display:"消耗"`                     // 消耗
	CpaBid               float64 `json:"cpa_bid" display:"出价"`                       // 出价
	RoiGoal              float64 `json:"roi_goal" display:"ROI目标"`                   // roi

	gorm.Model
}

func (AdReportItem) TableName() string {
	return "ad_report_item"
}

func (a *AdReportItem) Display() utils.SortedList {
	var ret []*utils.KV
	ret = append(ret, &utils.KV{Key: "id", Value: a.ID})
	value := reflect.ValueOf(a).Elem()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		key := value.Type().Field(i).Tag.Get("display")
		if key == "" {
			continue
		}
		ret = append(ret, &utils.KV{Key: key, Value: field.Interface()})
	}
	ret = append(ret, &utils.KV{Key: "created_at", Value: a.CreatedAt})
	ret = append(ret, &utils.KV{Key: "updated_at", Value: a.UpdatedAt})
	return ret
}
