package model

import "github.com/JingruiLea/ad_boost/model/ttypes"

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
