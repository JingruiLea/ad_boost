package model

import (
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"gorm.io/gorm"
)

type UniPromotionReport struct {
	UniAdID                      int64                `json:"uni_ad_id"`
	StartTime                    string               `json:"start_time"`
	EndTime                      string               `json:"end_time"`
	ModifyTime                   string               `json:"modify_time"`
	CreateTime                   string               `json:"create_time"`
	MarketingGoal                ttypes.MarketingGoal `json:"marketing_goal"`
	Roi2Goal                     float64              `json:"roi2_goal"`
	BudgetMode                   ttypes.BudgetMode    `json:"budget_mode"`
	Budget                       float64              `json:"budget"`
	Status                       ttypes.AdStatus      `json:"status"`
	OptStatus                    ttypes.OptStatus     `json:"opt_status"`
	DeliverySeconds              int                  `json:"delivery_seconds"`
	RoomInfo                     string               `json:"room_info"`
	StatCost                     float64              `json:"stat_cost"`
	TotalPrepayAndPayOrderRoi2   float64              `json:"total_prepay_and_pay_order_roi2"`
	TotalPayOrderGmvForRoi2      int                  `json:"total_pay_order_gmv_for_roi2"`
	TotalPayOrderCountForRoi2    float64              `json:"total_pay_order_count_for_roi2"`
	TotalCostPerPayOrderForRoi2  float64              `json:"total_cost_per_pay_order_for_roi2"`
	TotalPrepayOrderCountForRoi2 int                  `json:"total_prepay_order_count_for_roi2"`
	TotalPrepayOrderGmvForRoi2   float64              `json:"total_prepay_order_gmv_for_roi2"`

	gorm.Model
}
