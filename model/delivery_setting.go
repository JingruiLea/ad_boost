package model

import (
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
)

type DeliverySetting struct {
	SmartBidType          ttypes.SmartBidType       `json:"smart_bid_type"`           // 投放场景（出价方式）
	ExternalAction        ttypes.ExternalAction     `json:"external_action"`          // 转化目标
	DeepExternalAction    ttypes.DeepExternalAction `json:"deep_external_action"`     // 深度转化目标
	DeepBidType           ttypes.DeepBidType        `json:"deep_bid_type"`            // 深度出价方式
	ROIGoal               float64                   `json:"roi_goal"`                 // 支付ROI目标
	Budget                float64                   `json:"budget"`                   // 预算
	ReviveBudget          float64                   `json:"revive_budget"`            // 复活预算
	BudgetMode            ttypes.BudgetMode         `json:"budget_mode"`              // 预算类型
	CPABid                float64                   `json:"cpa_bid"`                  // 转化出价
	VideoScheduleType     ttypes.VideoScheduleType  `json:"video_schedule_type"`      // 短视频投放日期选择方式
	LiveScheduleType      ttypes.LiveScheduleType   `json:"live_schedule_type"`       // 直播间投放时段选择方式
	StartTime             string                    `json:"start_time"`               // 投放开始时间
	EndTime               string                    `json:"end_time"`                 // 投放结束时间
	ScheduleTime          string                    `json:"schedule_time"`            // 投放时段
	ScheduleFixedRange    int                       `json:"schedule_fixed_range"`     // 固定投放时长
	EnableAutoPause       int                       `json:"enable_auto_pause"`        // 是否启用超成本自动暂停
	AutoManageStrategyCmd int                       `json:"auto_manage_strategy_cmd"` // 托管策略
	EnableFollowMaterial  int                       `json:"enable_follow_material"`   // 是否优质素材自动同步投放
	ProductNewOpen        bool                      `json:"product_new_open"`         // 是否开启新品加速
	QCPXMode              ttypes.QcpxMode           `json:"qcpx_mode"`                // 智能优惠券状态
	AllowQCPX             bool                      `json:"allow_qcpx"`               // 是否支持智能优惠券
}

func (d *DeliverySetting) Display() utils.SortedList {
	return []*utils.KV{
		{Key: "投放场景（出价方式)", Value: d.SmartBidType},
		{Key: "转化目标", Value: d.ExternalAction},
		{Key: "深度转化目标", Value: d.DeepExternalAction},
		{Key: "深度出价方式", Value: d.DeepBidType},
		{Key: "支付ROI目标", Value: d.ROIGoal},
		{Key: "预算", Value: d.Budget},
		{Key: "复活预算", Value: d.ReviveBudget},
		{Key: "预算类型", Value: d.BudgetMode},
		{Key: "转化出价", Value: d.CPABid},
		{Key: "短视频投放日期选择方式", Value: d.VideoScheduleType},
		{Key: "直播间投放时段选择方式", Value: d.LiveScheduleType},
		{Key: "投放开始时间", Value: d.StartTime},
		{Key: "投放结束时间", Value: d.EndTime},
		{Key: "投放时段", Value: d.ScheduleTime},
		{Key: "固定投放时长", Value: d.ScheduleFixedRange},
		{Key: "是否开启新品加速", Value: d.ProductNewOpen},
		{Key: "智能优惠券状态，是否开启", Value: d.QCPXMode},
		{Key: "是否支持智能优惠券", Value: d.AllowQCPX},
	}
}
