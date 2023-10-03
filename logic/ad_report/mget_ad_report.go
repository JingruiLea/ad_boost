package ad_report

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
	"time"
)

func MGetAdReport(ctx context.Context, req *MGetAdReportReq) (*MGetAdReportRespData, error) {
	var resp MGetAdReportRespData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/report/ad/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "MGetAdReport httpclient.NewClient().Post error: %v", err)
		return nil, err
	}
	return &resp, nil
}

func MGetCommonAdReport(ctx context.Context, advertiserID int64, adIDs []int64) ([]*AdReport, error) {
	req := &MGetAdReportReq{
		AdvertiserID: advertiserID,
		StartDate:    time.Now().Format("2006-01-02"),
		EndDate:      time.Now().Format("2006-01-02"),
		Fields:       MGetAdReportFieldAllOrderCreateRoi7Days.Common(),
		Filtering: &MGetAdReportFiltering{
			MarketingGoal: ttypes.MarketingGoalLivePromGoods,
		},
		Page:     1,
		PageSize: len(adIDs),
	}
	req.StartDate = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	req.EndDate = time.Now().Format("2006-01-02")
	req.Filtering.AdIds = adIDs
	resp, err := MGetAdReport(ctx, req)
	if err != nil {
		logs.CtxErrorf(ctx, "MGetAdReport error: %v", err)
		return nil, err
	}
	return resp.List, nil
}

type MGetAdReportReq struct {
	AdvertiserID    int64                       `json:"advertiser_id,omitempty"`    // 广告主id
	StartDate       string                      `json:"start_date,omitempty"`       // 开始时间，格式 2021-04-05
	EndDate         string                      `json:"end_date,omitempty"`         // 结束时间，格式 2021-04-05
	TimeGranularity MGetAdReportTimeGranularity `json:"time_granularity,omitempty"` // 时间粒度
	Fields          []MGetAdReportField         `json:"fields,omitempty"`           // 需要查询的消耗指标
	Filtering       *MGetAdReportFiltering      `json:"filtering,omitempty"`        // 过滤条件
	Page            int                         `json:"page,omitempty"`             // 页码
	PageSize        int                         `json:"page_size,omitempty"`        // 页面大小
	OrderField      string                      `json:"order_field,omitempty"`      // 排序字段
	OrderType       string                      `json:"order_type,omitempty"`       // 排序方式
}

type MGetAdReportTimeGranularity string

const (
	//TIME_GRANULARITY_DAILY 按天
	MGetAdReportTimeGranularityDaily MGetAdReportTimeGranularity = "TIME_GRANULARITY_DAILY"
	//TIME_GRANULARITY_HOURLY 按小时
	MGetAdReportTimeGranularityHourly MGetAdReportTimeGranularity = "TIME_GRANULARITY_HOURLY"
)

type MGetAdReportFiltering struct {
	AdIds          []int64               `json:"ad_ids,omitempty"`          // 广告计划id列表，最多支持100个
	MarketingGoal  ttypes.MarketingGoal  `json:"marketing_goal,omitempty"`  // 营销目标
	OrderPlatform  string                `json:"order_platform,omitempty"`  // 下单平台
	MarketingScene ttypes.MarketingScene `json:"marketing_scene,omitempty"` // 广告类型
	CampaignScene  ttypes.CampaignScene  `json:"campaign_scene,omitempty"`  // 营销场景
	SmartBidType   ttypes.SmartBidType   `json:"smart_bid_type,omitempty"`  // 投放场景（投放方式）
	Status         string                `json:"status,omitempty"`          // 按计划状态过滤
}

// 报表指标
type MGetAdReportField string

const (
	//MGetAdReportFieldStatCost 消耗
	MGetAdReportFieldStatCost MGetAdReportField = "stat_cost"
	// MGetAdReportFieldShowCnt 展示次数
	MGetAdReportFieldShowCnt MGetAdReportField = "show_cnt"
	// MGetAdReportFieldCtr 点击率
	MGetAdReportFieldCtr MGetAdReportField = "ctr"
	// MGetAdReportFieldCpmPlatform 平均千次展示费用
	MGetAdReportFieldCpmPlatform MGetAdReportField = "cpm_platform"
	// MGetAdReportFieldClickCnt 点击次数
	MGetAdReportFieldClickCnt MGetAdReportField = "click_cnt"
	// MGetAdReportFieldPayOrderCount 直接成交订单数
	MGetAdReportFieldPayOrderCount MGetAdReportField = "pay_order_count"
	// MGetAdReportFieldPayOrderAmount 直接成交金额
	MGetAdReportFieldPayOrderAmount MGetAdReportField = "pay_order_amount"
	// MGetAdReportFieldPrepayAndPayOrderRoi 直接支付roi
	MGetAdReportFieldPrepayAndPayOrderRoi MGetAdReportField = "prepay_and_pay_order_roi"
	// MGetAdReportFieldCreateOrderCount 直接下单订单数
	MGetAdReportFieldCreateOrderCount MGetAdReportField = "create_order_count"
	// MGetAdReportFieldCreateOrderAmount 直接下单金额
	MGetAdReportFieldCreateOrderAmount MGetAdReportField = "create_order_amount"
	// MGetAdReportFieldCreateOrderRoi 直接下单roi
	MGetAdReportFieldCreateOrderRoi MGetAdReportField = "create_order_roi"
	// MGetAdReportFieldPrepayOrderCount 直接预售订单数
	MGetAdReportFieldPrepayOrderCount MGetAdReportField = "prepay_order_count"
	// MGetAdReportFieldPrepayOrderAmount 直接预售金额
	MGetAdReportFieldPrepayOrderAmount MGetAdReportField = "prepay_order_amount"
	// MGetAdReportFieldDyFollow 新增粉丝数
	MGetAdReportFieldDyFollow MGetAdReportField = "dy_follow"
	// MGetAdReportFieldConvertCnt 转化数
	MGetAdReportFieldConvertCnt MGetAdReportField = "convert_cnt"
	// MGetAdReportFieldConvertCost 转化成本
	MGetAdReportFieldConvertCost MGetAdReportField = "convert_cost"
	// MGetAdReportFieldConvertRate 转化率
	MGetAdReportFieldConvertRate MGetAdReportField = "convert_rate"
	// MGetAdReportFieldDyShare 分享次数
	MGetAdReportFieldDyShare MGetAdReportField = "dy_share"
	// MGetAdReportFieldDyComment 评论次数
	MGetAdReportFieldDyComment MGetAdReportField = "dy_comment"
	// MGetAdReportFieldDyLike 点赞次数
	MGetAdReportFieldDyLike MGetAdReportField = "dy_like"
	// MGetAdReportFieldLivePayOrderCostPerOrder 成交客单价
	MGetAdReportFieldLivePayOrderCostPerOrder MGetAdReportField = "live_pay_order_cost_per_order"
	// MGetAdReportFieldLubanLiveEnterCnt 直播间观看人次
	MGetAdReportFieldLubanLiveEnterCnt MGetAdReportField = "luban_live_enter_cnt"
	// MGetAdReportFieldLiveWatchOneMinuteCount 直播间超过1分钟观看人次
	MGetAdReportFieldLiveWatchOneMinuteCount MGetAdReportField = "live_watch_one_minute_count"
	// MGetAdReportFieldLiveFansClubJoinCnt 直播间新加团人次
	MGetAdReportFieldLiveFansClubJoinCnt MGetAdReportField = "live_fans_club_join_cnt"
	// MGetAdReportFieldLubanLiveSlidecartClickCnt 直播间查看购物车次数
	MGetAdReportFieldLubanLiveSlidecartClickCnt MGetAdReportField = "luban_live_slidecart_click_cnt"
	// MGetAdReportFieldLubanLiveClickProductCnt 直播间商品点击次数
	MGetAdReportFieldLubanLiveClickProductCnt MGetAdReportField = "luban_live_click_product_cnt"
	// MGetAdReportFieldLubanLiveCommentCnt 直播间评论次数
	MGetAdReportFieldLubanLiveCommentCnt MGetAdReportField = "luban_live_comment_cnt"
	// MGetAdReportFieldLubanLiveShareCnt 直播间分享次数
	MGetAdReportFieldLubanLiveShareCnt MGetAdReportField = "luban_live_share_cnt"
	// MGetAdReportFieldLubanLiveGiftCnt 直播间打赏次数
	MGetAdReportFieldLubanLiveGiftCnt MGetAdReportField = "luban_live_gift_cnt"
	// MGetAdReportFieldLubanLiveGiftAmount 直播间音浪收入
	MGetAdReportFieldLubanLiveGiftAmount MGetAdReportField = "luban_live_gift_amount"
	// MGetAdReportFieldTotalPlay 播放数
	MGetAdReportFieldTotalPlay MGetAdReportField = "total_play"
	// MGetAdReportFieldPlayDuration3s 3s播放数
	MGetAdReportFieldPlayDuration3s MGetAdReportField = "play_duration_3s"
	// MGetAdReportFieldPlay25FeedBreak 25%进度播放数
	MGetAdReportFieldPlay25FeedBreak MGetAdReportField = "play_25_feed_break"
	// MGetAdReportFieldPlay50FeedBreak 50%进度播放数
	MGetAdReportFieldPlay50FeedBreak MGetAdReportField = "play_50_feed_break"
	// MGetAdReportFieldPlay75FeedBreak 75%进度播放数
	MGetAdReportFieldPlay75FeedBreak MGetAdReportField = "play_75_feed_break"
	// MGetAdReportFieldPlayOver 播放完成数
	MGetAdReportFieldPlayOver MGetAdReportField = "play_over"
	// MGetAdReportFieldPlayOverRate 完播率
	MGetAdReportFieldPlayOverRate MGetAdReportField = "play_over_rate"
	// MGetAdReportFieldCpcPlatform 平均点击单价
	MGetAdReportFieldCpcPlatform MGetAdReportField = "cpc_platform"
	// MGetAdReportFieldDeepConvertCnt 深度转化次数
	MGetAdReportFieldDeepConvertCnt MGetAdReportField = "deep_convert_cnt"
	// MGetAdReportFieldDeepConvertCost 深度转化成本
	MGetAdReportFieldDeepConvertCost MGetAdReportField = "deep_convert_cost"
	// MGetAdReportFieldDeepConvertRate 深度转化率
	MGetAdReportFieldDeepConvertRate MGetAdReportField = "deep_convert_rate"
	// MGetAdReportFieldAttributionConvertCnt 转化数(计费时间)
	MGetAdReportFieldAttributionConvertCnt MGetAdReportField = "attribution_convert_cnt"
	// MGetAdReportFieldAttributionConvertRate 转化率(计费时间)
	MGetAdReportFieldAttributionConvertRate MGetAdReportField = "attribution_convert_rate"
	// MGetAdReportFieldAttributionConvertCost 转化成本(计费时间)
	MGetAdReportFieldAttributionConvertCost MGetAdReportField = "attribution_convert_cost"
	// MGetAdReportFieldAttributionDeepConvertCnt 深度转化次数(计费时间)
	MGetAdReportFieldAttributionDeepConvertCnt MGetAdReportField = "attribution_deep_convert_cnt"
	// MGetAdReportFieldAttributionDeepConvertCost 深度转化成本(计费时间)
	MGetAdReportFieldAttributionDeepConvertCost MGetAdReportField = "attribution_deep_convert_cost"
	// MGetAdReportFieldAttributionDeepConvertRate 深度转化率(计费时间)
	MGetAdReportFieldAttributionDeepConvertRate MGetAdReportField = "attribution_deep_convert_rate"
	// MGetAdReportFieldAllOrderCreateRoi7Days 7日总下单ROI
	MGetAdReportFieldAllOrderCreateRoi7Days MGetAdReportField = "all_order_create_roi_7days"
	// MGetAdReportFieldAllOrderPayRoi7Days 7日总支付ROI
	MGetAdReportFieldAllOrderPayRoi7Days MGetAdReportField = "all_order_pay_roi_7days"
	// MGetAdReportFieldAllOrderPayCount7Days 7日总成交订单
	MGetAdReportFieldAllOrderPayCount7Days MGetAdReportField = "all_order_pay_count_7days"
	// MGetAdReportFieldAllOrderPayGmv7Days 7日总成交金额
	MGetAdReportFieldAllOrderPayGmv7Days MGetAdReportField = "all_order_pay_gmv_7days"
	// MGetAdReportFieldPayOrderCostPerOrder 直接成交客单价
	MGetAdReportFieldPayOrderCostPerOrder MGetAdReportField = "pay_order_cost_per_order"
	// MGetAdReportFieldCreateOrderCouponAmount 下单智能优惠券金额
	MGetAdReportFieldCreateOrderCouponAmount MGetAdReportField = "create_order_coupon_amount"
	// MGetAdReportFieldPayOrderCouponAmount 成交智能优惠券金额
	MGetAdReportFieldPayOrderCouponAmount MGetAdReportField = "pay_order_coupon_amount"
	// MGetAdReportFieldIndirectOrderCreateCount7Days 间接下单订单数
	MGetAdReportFieldIndirectOrderCreateCount7Days MGetAdReportField = "indirect_order_create_count_7days"
	// MGetAdReportFieldIndirectOrderCreateGmv7Days 间接下单金额
	MGetAdReportFieldIndirectOrderCreateGmv7Days MGetAdReportField = "indirect_order_create_gmv_7days"
	// MGetAdReportFieldIndirectOrderPayCount7Days 间接成交订单数
	MGetAdReportFieldIndirectOrderPayCount7Days MGetAdReportField = "indirect_order_pay_count_7days"
	// MGetAdReportFieldIndirectOrderPayGmv7Days 间接成交金额
	MGetAdReportFieldIndirectOrderPayGmv7Days MGetAdReportField = "indirect_order_pay_gmv_7days"
	// MGetAdReportFieldIndirectOrderPrepayCount7Days 间接预售订单数
	MGetAdReportFieldIndirectOrderPrepayCount7Days MGetAdReportField = "indirect_order_prepay_count_7days"
	// MGetAdReportFieldIndirectOrderPrepayGmv7Days 间接预售金额
	MGetAdReportFieldIndirectOrderPrepayGmv7Days MGetAdReportField = "indirect_order_prepay_gmv_7days"
	// MGetAdReportFieldUpliftPayOrderCount 增效成交订单数 (短视频带货不支持)
	MGetAdReportFieldUpliftPayOrderCount MGetAdReportField = "uplift_pay_order_count"
	// MGetAdReportFieldUpliftPayOrderGmv 增效成交金额 (短视频带货不支持)
	MGetAdReportFieldUpliftPayOrderGmv MGetAdReportField = "uplift_pay_order_gmv"
	// MGetAdReportFieldUpliftPayOrderRoi 增效支付ROI (短视频带货不支持)
	MGetAdReportFieldUpliftPayOrderRoi MGetAdReportField = "uplift_pay_order_roi"
	// MGetAdReportFieldUpliftCreateOrderCount 增效下单订单数 (短视频带货不支持)
	MGetAdReportFieldUpliftCreateOrderCount MGetAdReportField = "uplift_create_order_count"
	// MGetAdReportFieldUpliftCreateOrderGmv 增效下单金额 (短视频带货不支持)
	MGetAdReportFieldUpliftCreateOrderGmv MGetAdReportField = "uplift_create_order_gmv"
	// MGetAdReportFieldUpliftCreateOrderRoi 增效下单ROI (短视频带货不支持)
	MGetAdReportFieldUpliftCreateOrderRoi MGetAdReportField = "uplift_create_order_roi"
	// MGetAdReportFieldQianchuanFirstOrderCnt 首单新客人数
	MGetAdReportFieldQianchuanFirstOrderCnt MGetAdReportField = "qianchuan_first_order_cnt"
	// MGetAdReportFieldQianchuanFirstOrderRate 首单新客订单占比
	MGetAdReportFieldQianchuanFirstOrderRate MGetAdReportField = "qianchuan_first_order_rate"
	// MGetAdReportFieldQianchuanFirstOrderConvertCost 首单新客转化成本
	MGetAdReportFieldQianchuanFirstOrderConvertCost MGetAdReportField = "qianchuan_first_order_convert_cost"
	// MGetAdReportFieldQianchuanFirstOrderDirectPayGmv 首单新客直接成交金额
	MGetAdReportFieldQianchuanFirstOrderDirectPayGmv MGetAdReportField = "qianchuan_first_order_direct_pay_gmv"
	// MGetAdReportFieldQianchuanFirstOrderDirectPayOrderRoi 首单新客直接支付ROI
	MGetAdReportFieldQianchuanFirstOrderDirectPayOrderRoi MGetAdReportField = "qianchuan_first_order_direct_pay_order_roi"
	// MGetAdReportFieldQianchuanFirstOrderLtv30 首单新客30天累计成交金额
	MGetAdReportFieldQianchuanFirstOrderLtv30 MGetAdReportField = "qianchuan_first_order_ltv30"
	// MGetAdReportFieldQianchuanFirstOrderRoi30 首单新客30天累计支付ROI
	MGetAdReportFieldQianchuanFirstOrderRoi30 MGetAdReportField = "qianchuan_first_order_roi30"
)

func (MGetAdReportField) Common() []MGetAdReportField {
	return []MGetAdReportField{
		MGetAdReportFieldStatCost,
		MGetAdReportFieldShowCnt,
		MGetAdReportFieldCtr,
		MGetAdReportFieldCpmPlatform,
		MGetAdReportFieldClickCnt,
		MGetAdReportFieldPayOrderCount,
		MGetAdReportFieldPayOrderAmount,
		MGetAdReportFieldPrepayAndPayOrderRoi,
		MGetAdReportFieldDyFollow,
		MGetAdReportFieldConvertCnt,
		MGetAdReportFieldConvertCost,
		MGetAdReportFieldConvertRate,
	}
}

type MGetAdReportResp struct {
	Data *MGetAdReportRespData `json:"data"`
	ttypes.BaseResp
}

type MGetAdReportRespData struct {
	List     []*AdReport     `json:"list"`
	PageInfo ttypes.PageInfo `json:"page_info"`
}

type AdReport struct {
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
}

func (a *AdReport) ToModel(cpa, roi float64) *model.AdReportItem {
	ret := &model.AdReportItem{
		AdID:                 a.AdID,
		AdvertiserId:         a.AdvertiserId,
		ClickCnt:             a.ClickCnt,
		ConvertCnt:           a.ConvertCnt,
		ConvertCost:          a.ConvertCost,
		ConvertRate:          a.ConvertRate,
		CpmPlatform:          a.CpmPlatform,
		Ctr:                  a.Ctr,
		DyFollow:             a.DyFollow,
		PayOrderAmount:       a.PayOrderAmount,
		PayOrderCount:        a.PayOrderCount,
		PrepayAndPayOrderRoi: a.PrepayAndPayOrderRoi,
		ShowCnt:              a.ShowCnt,
		StatCost:             a.StatCost,
		CpaBid:               cpa,
		RoiGoal:              roi,
	}
	return ret
}
