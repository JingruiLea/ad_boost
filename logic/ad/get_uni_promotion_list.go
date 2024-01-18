package ad

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetUniPromotionList(ctx context.Context, req *GetUniPromotionListReq) (*GetUniPromotionListRespData, error) {
	var resp GetUniPromotionListRespData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/uni_promotion/list/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetUniPromotionList httpclient.NewClient().AdGet error: %v", err)
		return nil, err
	}
	return &resp, nil
}

type GetUniPromotionListReq struct {
	// AdvertiserID 表示广告主ID。
	AdvertiserID int64 `json:"advertiser_id"`
	// StartTime 表示广告开始时间。格式: "YYYY-MM-DD HH:mm:ss"。
	StartTime string `json:"start_time"`
	// EndTime 表示广告结束时间。 格式: "YYYY-MM-DD HH:mm:ss"。
	EndTime string `json:"end_time"`
	// MarketingGoal 表示按营销目标过滤。
	MarketingGoal ttypes.MarketingGoal `json:"marketing_goal"`
	// Fields 表示需要查询的消耗指标。
	Fields []UniPromotionStatField `json:"fields"`
	// OrderType 表示排序方式，可选字段。
	OrderType *ttypes.OrderType `json:"order_type,omitempty"`
	// OrderField 表示排序字段，可选字段。
	OrderField *string `json:"order_field,omitempty"`
	// Page 表示页码，可选字段。
	Page *int `json:"page,omitempty"`
	// PageSize 表示页面大小，可选字段。
	PageSize *int `json:"page_size,omitempty"`
}

// Data 表示JSON返回值的顶层结构。
type GetUniPromotionListRespData struct {
	AdList   []*UniAdInfo    `json:"ad_list"`   // 全域推广列表
	PageInfo ttypes.PageInfo `json:"page_info"` // 分页信息
}

// UniAdInfo 表示全域广告信息。
type UniAdInfo struct {
	ID              int64                `json:"id"`               // 推广id
	StartTime       string               `json:"start_time"`       // 当前周期开始时间
	EndTime         string               `json:"end_time"`         // 当前周期结束时间
	ModifyTime      string               `json:"modify_time"`      // 修改时间
	CreateTime      string               `json:"create_time"`      // 创建时间
	MarketingGoal   ttypes.MarketingGoal `json:"marketing_goal"`   // 营销目标
	Roi2Goal        float64              `json:"roi2_goal"`        // 支付ROI目标
	BudgetMode      ttypes.BudgetMode    `json:"budget_mode"`      // 预算类型
	Budget          float64              `json:"budget"`           // 预算
	Status          ttypes.AdStatus      `json:"status"`           // 投放状态
	OptStatus       ttypes.OptStatus     `json:"opt_status"`       // 操作状态
	DeliverySeconds int                  `json:"delivery_seconds"` // 投放时长
	RoomInfo        []RoomInfo           `json:"room_info"`        // 主播信息
	StatsInfo       StatsInfo            `json:"stats_info"`       // 消耗指标
}

func (u UniAdInfo) ToModel() *model.UniPromotionReport {
	ret := &model.UniPromotionReport{
		UniAdID:                      u.ID,
		StartTime:                    u.StartTime,
		EndTime:                      u.EndTime,
		ModifyTime:                   u.ModifyTime,
		CreateTime:                   u.CreateTime,
		MarketingGoal:                u.MarketingGoal,
		Roi2Goal:                     u.Roi2Goal,
		BudgetMode:                   u.BudgetMode,
		Budget:                       u.Budget,
		Status:                       u.Status,
		OptStatus:                    u.OptStatus,
		DeliverySeconds:              u.DeliverySeconds,
		RoomInfo:                     utils.GetJsonStr(u.RoomInfo),
		StatCost:                     u.StatsInfo.StatCost,
		TotalPrepayAndPayOrderRoi2:   u.StatsInfo.TotalPrepayAndPayOrderRoi2,
		TotalPayOrderGmvForRoi2:      u.StatsInfo.TotalPayOrderGmvForRoi2,
		TotalPayOrderCountForRoi2:    u.StatsInfo.TotalPayOrderCountForRoi2,
		TotalCostPerPayOrderForRoi2:  u.StatsInfo.TotalCostPerPayOrderForRoi2,
		TotalPrepayOrderCountForRoi2: u.StatsInfo.TotalPrepayOrderCountForRoi2,
		TotalPrepayOrderGmvForRoi2:   u.StatsInfo.TotalPrepayOrderGmvForRoi2,
	}
	return ret
}

// RoomInfo 表示主播信息。
type RoomInfo struct {
	AnchorID     string `json:"anchor_id"`     // 主播ID
	AnchorName   string `json:"anchor_name"`   // 主播名称
	AnchorAvatar string `json:"anchor_avatar"` // 主播头像
}

// StatsInfo 表示消耗指标。
type StatsInfo struct {
	StatCost                     float64 `json:"stat_cost"`                         // 整体消耗
	TotalPrepayAndPayOrderRoi2   float64 `json:"total_prepay_and_pay_order_roi2"`   // 整体支付ROI
	TotalPayOrderGmvForRoi2      int     `json:"total_pay_order_gmv_for_roi2"`      // 整体成交金额
	TotalPayOrderCountForRoi2    float64 `json:"total_pay_order_count_for_roi2"`    // 整体成交订单数
	TotalCostPerPayOrderForRoi2  float64 `json:"total_cost_per_pay_order_for_roi2"` // 整体成交订单成本
	TotalPrepayOrderCountForRoi2 int     `json:"total_prepay_order_count_for_roi2"` // 整体预售订单数
	TotalPrepayOrderGmvForRoi2   float64 `json:"total_prepay_order_gmv_for_roi2"`   // 整体预售订单金额
}

type UniPromotionStatField string

const (
	UniPromotionStatFieldStatCost                     UniPromotionStatField = "stat_cost"                         // 整体消耗
	UniPromotionStatFieldTotalPrepayAndPayOrderRoi2   UniPromotionStatField = "total_prepay_and_pay_order_roi2"   // 整体支付ROI
	UniPromotionStatFieldTotalPayOrderGmvForRoi2      UniPromotionStatField = "total_pay_order_gmv_for_roi2"      // 整体成交金额
	UniPromotionStatFieldTotalPayOrderCountForRoi2    UniPromotionStatField = "total_pay_order_count_for_roi2"    // 整体成交订单数
	UniPromotionStatFieldTotalCostPerPayOrderForRoi2  UniPromotionStatField = "total_cost_per_pay_order_for_roi2" // 整体成交订单成本
	UniPromotionStatFieldTotalPrepayOrderCountForRoi2 UniPromotionStatField = "total_prepay_order_count_for_roi2" // 整体预售订单数
	UniPromotionStatFieldTotalPrepayOrderGmvForRoi2   UniPromotionStatField = "total_prepay_order_gmv_for_roi2"   // 整体预售订单金额
)

func (s UniPromotionStatField) All() []UniPromotionStatField {
	return []UniPromotionStatField{
		UniPromotionStatFieldStatCost,
		UniPromotionStatFieldTotalPrepayAndPayOrderRoi2,
		UniPromotionStatFieldTotalPayOrderGmvForRoi2,
		UniPromotionStatFieldTotalPayOrderCountForRoi2,
		UniPromotionStatFieldTotalCostPerPayOrderForRoi2,
		UniPromotionStatFieldTotalPrepayOrderCountForRoi2,
		UniPromotionStatFieldTotalPrepayOrderGmvForRoi2,
	}
}
