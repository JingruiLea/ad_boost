package model

import (
	"gorm.io/gorm"
)

type AdReportItem struct {
	ID                   int64   `json:"id"`                       // 自增id
	AdID                 int64   `json:"ad_id"`                    // 广告计划id
	AdvertiserId         int64   `json:"advertiser_id"`            // 广告主id
	ClickCnt             int     `json:"click_cnt"`                // 点击次数
	ConvertCnt           int     `json:"convert_cnt"`              // 转化数
	ConvertCost          float64 `json:"convert_cost"`             // 转化成本
	ConvertRate          float64 `json:"convert_rate"`             // 转化率
	CpmPlatform          float64 `json:"cpm_platform"`             // 平均千次展示费用
	Ctr                  float64 `json:"ctr"`                      // 点击率
	DyFollow             int     `json:"dy_follow"`                // 新增粉丝数
	PayOrderAmount       float64 `json:"pay_order_amount"`         // 直接成交金额
	PayOrderCount        int     `json:"pay_order_count"`          // 直接成交订单数
	PrepayAndPayOrderRoi float64 `json:"prepay_and_pay_order_roi"` // 直接支付roi
	ShowCnt              int     `json:"show_cnt"`                 // 展示次数
	StatCost             float64 `json:"stat_cost"`                // 消耗
	CpaBid               float64 `json:"cpa_bid"`                  // 出价
	RoiGoal              float64 `json:"roi_goal"`                 // roi

	gorm.Model
}

func (AdReportItem) TableName() string {
	return "ad_report_item"
}
