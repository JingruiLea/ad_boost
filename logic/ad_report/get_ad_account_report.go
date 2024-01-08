package ad_report

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetAdAccountReport(ctx context.Context, req *GetAdAccountReportReq) (*AdvertiserReportRespData, error) {
	var resp AdvertiserReportRespData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/report/advertiser/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "MGetAdvertiserReport httpclient.NewClient().Post error: %v", err)
		return nil, err
	}
	logs.CtxInfof(ctx, "MGetAdvertiserReport respMap: %s", utils.GetJsonStr(resp))
	return &resp, nil
}

type GetAdAccountReportReq struct {
	AdvertiserID    int64                     `json:"advertiser_id"`              // 千川广告主账户id
	StartDate       string                    `json:"start_date"`                 // 开始时间，格式 2021-04-05
	EndDate         string                    `json:"end_date"`                   // 结束时间，格式 2021-04-05
	TimeGranularity *ttypes.TimeGranularity   `json:"time_granularity,omitempty"` // 时间粒度
	Fields          []GetAdAccountReportField `json:"fields"`                     // 需要查询的消耗指标
	Filtering       GetAdAccountReportFilter  `json:"filtering"`                  // 过滤条件
	OrderField      *string                   `json:"order_field,omitempty"`      // 排序字段
	OrderType       *string                   `json:"order_type,omitempty"`       // 排序方式
	Page            *int                      `json:"page,omitempty"`             // 页码，默认为1
	PageSize        *int                      `json:"page_size,omitempty"`        // 页面大小，默认为10
}

type GetAdAccountReportFilter struct {
	MarketingGoal  ttypes.MarketingGoal   `json:"marketing_goal"`            // 营销目标
	OrderPlatform  *ttypes.OrderPlatform  `json:"order_platform,omitempty"`  // 下单平台
	MarketingScene *ttypes.MarketingScene `json:"marketing_scene,omitempty"` // 广告类型
	CampaignScene  *ttypes.CampaignScene  `json:"campaign_scene,omitempty"`  // 营销场景
	SmartBidType   *ttypes.SmartBidType   `json:"smart_bid_type,omitempty"`  // 投放场景（投放方式）
	Status         *ttypes.AdStatus       `json:"status,omitempty"`          // 按计划状态过滤
	AwemeIDs       []int                  `json:"aweme_ids,omitempty"`       // 按抖音id过滤
}

type AdvertiserReportRespData struct {
	AdvertiserID                               int     `json:"advertiser_id"`                                     // 广告主id
	StatDatetime                               string  `json:"stat_datetime"`                                     // 数据起始时间
	StatCost                                   float64 `json:"stat_cost"`                                         // 消耗
	ShowCnt                                    int     `json:"show_cnt"`                                          // 展示次数
	Ctr                                        float64 `json:"ctr"`                                               // 点击率
	CpmPlatform                                float64 `json:"cpm_platform"`                                      // 平均千次展示费用
	ClickCnt                                   int     `json:"click_cnt"`                                         // 点击次数
	PayOrderCount                              int     `json:"pay_order_count"`                                   // 直接成交订单数
	PayOrderAmount                             float64 `json:"pay_order_amount"`                                  // 直接成交金额
	PrepayAndPayOrderROI                       float64 `json:"prepay_and_pay_order_roi"`                          // 直接支付roi
	CreateOrderCount                           int     `json:"create_order_count"`                                // 直接下单订单数
	CreateOrderAmount                          float64 `json:"create_order_amount"`                               // 直接下单金额
	CreateOrderROI                             float64 `json:"create_order_roi"`                                  // 直接下单roi
	PrepayOrderCount                           int     `json:"prepay_order_count"`                                // 直接预售订单数
	PrepayOrderAmount                          float64 `json:"prepay_order_amount"`                               // 直接预售金额
	DyFollow                                   int     `json:"dy_follow"`                                         // 新增粉丝数
	TotalPlay                                  int     `json:"total_play"`                                        // 播放数
	PlayDuration3s                             int     `json:"play_duration_3s"`                                  // 3s播放数
	Play25FeedBreak                            int     `json:"play_25_feed_break"`                                // 25%进度播放数
	Play50FeedBreak                            int     `json:"play_50_feed_break"`                                // 50%进度播放数
	Play75FeedBreak                            int     `json:"play_75_feed_break"`                                // 75%进度播放数
	PlayOver                                   int     `json:"play_over"`                                         // 播放完成数
	PlayOverRate                               float64 `json:"play_over_rate"`                                    // 完播率
	CpcPlatform                                float64 `json:"cpc_platform"`                                      // 平均点击单价
	DeepConvertCnt                             int     `json:"deep_convert_cnt"`                                  // 深度转化次数
	DeepConvertCost                            float64 `json:"deep_convert_cost"`                                 // 深度转化成本
	DeepConvertRate                            float64 `json:"deep_convert_rate"`                                 // 深度转化率
	AttributionConvertCnt                      int     `json:"attribution_convert_cnt"`                           // 转化数（计费时间）
	AttributionConvertRate                     float64 `json:"attribution_convert_rate"`                          // 转化率（计费时间）
	AttributionConvertCost                     float64 `json:"attribution_convert_cost"`                          // 转化成本（计费时间）
	AttributionDeepConvertCnt                  int     `json:"attribution_deep_convert_cnt"`                      // 深度转化次数（计费时间）
	AttributionDeepConvertCost                 float64 `json:"attribution_deep_convert_cost"`                     // 深度转化成本（计费时间）
	AttributionDeepConvertRate                 float64 `json:"attribution_deep_convert_rate"`                     // 深度转化率（计费时间）
	AllOrderCreateROI7days                     float64 `json:"all_order_create_roi_7days"`                        // 7日总下单ROI
	AllOrderPayROI7days                        float64 `json:"all_order_pay_roi_7days"`                           // 7日总支付ROI
	AllOrderPayCount7days                      int     `json:"all_order_pay_count_7days"`                         // 7日总成交订单
	AllOrderPayGMV7days                        float64 `json:"all_order_pay_gmv_7days"`                           // 7日总成交金额
	PayOrderCostPerOrder                       float64 `json:"pay_order_cost_per_order"`                          // 直接成交客单价
	CreateOrderCouponAmount                    float64 `json:"create_order_coupon_amount"`                        // 下单智能优惠券金额
	PayOrderCouponAmount                       float64 `json:"pay_order_coupon_amount"`                           // 成交智能优惠券金额
	IndirectOrderCreateCount7days              int     `json:"indirect_order_create_count_7days"`                 // 间接下单订单数
	IndirectOrderCreateGMV7days                float64 `json:"indirect_order_create_gmv_7days"`                   // 间接下单金额
	IndirectOrderPayCount7days                 int     `json:"indirect_order_pay_count_7days"`                    // 间接成交订单数
	IndirectOrderPayGMV7days                   float64 `json:"indirect_order_pay_gmv_7days"`                      // 间接成交金额
	IndirectOrderPrepayCount7days              int     `json:"indirect_order_prepay_count_7days"`                 // 间接预售订单数
	IndirectOrderPrepayGMV7days                float64 `json:"indirect_order_prepay_gmv_7days"`                   // 间接预售金额
	QianchuanFirstOrderCnt                     int     `json:"qianchuan_first_order_cnt"`                         // 店铺首单新客人数
	QianchuanFirstOrderRate                    float64 `json:"qianchuan_first_order_rate"`                        // 店铺首单新客订单占比
	QianchuanFirstOrderConvertCost             float64 `json:"qianchuan_first_order_convert_cost"`                // 店铺首单新客转化成本
	QianchuanFirstOrderDirectPayGMV            float64 `json:"qianchuan_first_order_direct_pay_gmv"`              // 店铺首单新客直接成交金额
	QianchuanFirstOrderDirectPayOrderROI       float64 `json:"qianchuan_first_order_direct_pay_order_roi"`        // 店铺首单新客直接支付ROI
	QianchuanFirstOrderLTV30                   float64 `json:"qianchuan_first_order_ltv30"`                       // 店铺首单新客30天累计成交金额
	QianchuanFirstOrderROI30                   float64 `json:"qianchuan_first_order_roi30"`                       // 店铺首单新客30天累计支付ROI
	QianchuanBrandFirstOrderRate               float64 `json:"qianchuan_brand_first_order_rate"`                  // 品牌首单新客订单占比
	QianchuanBrandFirstOrderConvertCost        float64 `json:"qianchuan_brand_first_order_convert_cost"`          // 品牌首单新客转化成本
	QianchuanBrandFirstOrderLTV30              float64 `json:"qianchuan_brand_first_order_ltv30"`                 // 品牌首单新客30天累计成交金额
	QianchuanBrandFirstOrderDirectPayOrderROI  float64 `json:"qianchuan_brand_first_order_direct_pay_order_roi"`  // 品牌首单新客直接支付ROI
	QianchuanBrandFirstOrderROI30              float64 `json:"qianchuan_brand_first_order_roi30"`                 // 品牌首单新客30天累计支付ROI
	QianchuanBrandFirstOrderCnt                int     `json:"qianchuan_brand_first_order_cnt"`                   // 品牌首单新客数
	QianchuanBrandFirstOrderDirectPayGMV       float64 `json:"qianchuan_brand_first_order_direct_pay_gmv"`        // 品牌首单新客直接成交金额
	QianchuanAuthorFirstOrderRate              float64 `json:"qianchuan_author_first_order_rate"`                 // 抖音号首单新客订单占比
	QianchuanAuthorFirstOrderROI30             float64 `json:"qianchuan_author_first_order_roi30"`                // 抖音号首单新客30天累计支付ROI
	QianchuanAuthorFirstOrderCnt               int     `json:"qianchuan_author_first_order_cnt"`                  // 抖音号首单新客数
	QianchuanAuthorFirstOrderConvertCost       float64 `json:"qianchuan_author_first_order_convert_cost"`         // 抖音号首单新客转化成本
	QianchuanAuthorFirstOrderDirectPayGMV      float64 `json:"qianchuan_author_first_order_direct_pay_gmv"`       // 抖音号首单新客直接成交金额
	QianchuanAuthorFirstOrderDirectPayOrderROI float64 `json:"qianchuan_author_first_order_direct_pay_order_roi"` // 抖音号首单新客直接支付ROI
	QianchuanAuthorFirstOrderLTV30             float64 `json:"qianchuan_author_first_order_ltv30"`                // 抖音号首单新客30天累计成交金额
	LiveOrderSettleAmount14D                   float64 `json:"live_order_settle_amount_14d"`                      // 14天结算金额
	LiveOrderSettleCount14D                    int     `json:"live_order_settle_count_14d"`                       // 14天结算订单数
	LiveOrderSettleCountRate14D                float64 `json:"live_order_settle_count_rate_14d"`                  // 14天结算订单率
	AdLiveOrderSettleROI14D                    float64 `json:"ad_live_order_settle_roi_14d"`                      // 14天结算ROI
	AdLiveOrderSettleCost14D                   float64 `json:"ad_live_order_settle_cost_14d"`                     // 14天结算成本
	UnfinishedEstimateOrderGMV                 float64 `json:"unfinished_estimate_order_gmv"`                     // 未完结直接预售订单预估金额
	IndirectOrderUnfinishedEstimateGMV7Days    float64 `json:"indirect_order_unfinished_estimate_gmv_7days"`      // 未完结间接预售订单预估金额

	PageInfo ttypes.PageInfo `json:"page_info"` // 分页信息
}

func (f GetAdAccountReportField) Common() []GetAdAccountReportField {
	return []GetAdAccountReportField{
		GetAdAccountReportFieldStatCost,
		GetAdAccountReportFieldShowCnt,
		GetAdAccountReportFieldCtr,
		GetAdAccountReportFieldCpmPlatform,
		GetAdAccountReportFieldClickCnt,
		GetAdAccountReportFieldPayOrderCount,
		GetAdAccountReportFieldPayOrderAmount,
		GetAdAccountReportFieldPrepayAndPayOrderROI,
		GetAdAccountReportFieldCreateOrderCount,
		GetAdAccountReportFieldCreateOrderAmount,
		GetAdAccountReportFieldCreateOrderROI,
		GetAdAccountReportFieldPrepayOrderCount,
		GetAdAccountReportFieldPrepayOrderAmount,
		GetAdAccountReportFieldDyFollow,
		GetAdAccountReportFieldTotalPlay,
	}
}

func (f GetAdAccountReportField) All() []GetAdAccountReportField {
	return []GetAdAccountReportField{
		GetAdAccountReportFieldStatCost,
		GetAdAccountReportFieldShowCnt,
		GetAdAccountReportFieldCtr,
		GetAdAccountReportFieldCpmPlatform,
		GetAdAccountReportFieldClickCnt,
		GetAdAccountReportFieldPayOrderCount,
		GetAdAccountReportFieldPayOrderAmount,
		GetAdAccountReportFieldPrepayAndPayOrderROI,
		GetAdAccountReportFieldCreateOrderCount,
		GetAdAccountReportFieldCreateOrderAmount,
		GetAdAccountReportFieldCreateOrderROI,
		GetAdAccountReportFieldPrepayOrderCount,
		GetAdAccountReportFieldPrepayOrderAmount,
		GetAdAccountReportFieldDyFollow,
		GetAdAccountReportFieldTotalPlay,
		GetAdAccountReportFieldPlayDuration3s,
		GetAdAccountReportFieldPlay25FeedBreak,
		GetAdAccountReportFieldPlay50FeedBreak,
		GetAdAccountReportFieldPlay75FeedBreak,
		GetAdAccountReportFieldPlayOver,
		GetAdAccountReportFieldPlayOverRate,
		GetAdAccountReportFieldCpcPlatform,
		GetAdAccountReportFieldDeepConvertCnt,
		GetAdAccountReportFieldDeepConvertCost,
		GetAdAccountReportFieldDeepConvertRate,
		GetAdAccountReportFieldAttributionConvertCnt,
		GetAdAccountReportFieldAttributionConvertRate,
		GetAdAccountReportFieldAttributionConvertCost,
		GetAdAccountReportFieldAttributionDeepConvertCnt,
		GetAdAccountReportFieldAttributionDeepConvertCost,
		GetAdAccountReportFieldAttributionDeepConvertRate,
		GetAdAccountReportFieldAllOrderCreateROI7days,
		GetAdAccountReportFieldAllOrderPayROI7days,
		GetAdAccountReportFieldAllOrderPayCount7days,
		GetAdAccountReportFieldAllOrderPayGMV7days,
		GetAdAccountReportFieldPayOrderCostPerOrder,
		GetAdAccountReportFieldCreateOrderCouponAmount,
		GetAdAccountReportFieldPayOrderCouponAmount,
		GetAdAccountReportFieldIndirectOrderCreateCount7days,
		GetAdAccountReportFieldIndirectOrderCreateGMV7days,
		GetAdAccountReportFieldIndirectOrderPayCount7days,
		GetAdAccountReportFieldIndirectOrderPayGMV7days,
		GetAdAccountReportFieldIndirectOrderPrepayCount7days,
		GetAdAccountReportFieldIndirectOrderPrepayGMV7days,
		GetAdAccountReportFieldQianchuanFirstOrderCnt,
		GetAdAccountReportFieldQianchuanFirstOrderRate,
		GetAdAccountReportFieldQianchuanFirstOrderConvertCost,
		GetAdAccountReportFieldQianchuanFirstOrderDirectPayGMV,
		GetAdAccountReportFieldQianchuanFirstOrderDirectPayOrderROI,
		GetAdAccountReportFieldQianchuanFirstOrderLTV30,
		GetAdAccountReportFieldQianchuanFirstOrderROI30,
		GetAdAccountReportFieldQianchuanBrandFirstOrderRate,
		GetAdAccountReportFieldQianchuanBrandFirstOrderConvertCost,
		GetAdAccountReportFieldQianchuanBrandFirstOrderLTV30,
		GetAdAccountReportFieldQianchuanBrandFirstOrderDirectPayOrderROI,
		GetAdAccountReportFieldQianchuanBrandFirstOrderROI30,
		GetAdAccountReportFieldQianchuanBrandFirstOrderCnt,
		GetAdAccountReportFieldQianchuanBrandFirstOrderDirectPayGMV,
		GetAdAccountReportFieldQianchuanAuthorFirstOrderRate,
		GetAdAccountReportFieldQianchuanAuthorFirstOrderROI30,
		GetAdAccountReportFieldQianchuanAuthorFirstOrderCnt,
		GetAdAccountReportFieldQianchuanAuthorFirstOrderConvertCost,
		GetAdAccountReportFieldQianchuanAuthorFirstOrderDirectPayGMV,
		GetAdAccountReportFieldQianchuanAuthorFirstOrderDirectPayOrderROI,
		GetAdAccountReportFieldQianchuanAuthorFirstOrderLTV30,
		GetAdAccountReportFieldLiveOrderSettleAmount14D,
		GetAdAccountReportFieldLiveOrderSettleCount14D,
		GetAdAccountReportFieldLiveOrderSettleCountRate14D,
		GetAdAccountReportFieldAdLiveOrderSettleROI14D,
		GetAdAccountReportFieldAdLiveOrderSettleCost14D,
		GetAdAccountReportFieldUnfinishedEstimateOrderGMV,
		GetAdAccountReportFieldIndirectOrderUnfinishedEstimateGMV7Days,
	}
}

type GetAdAccountReportField string

const (
	GetAdAccountReportFieldStatCost                                   GetAdAccountReportField = "stat_cost"                                         // 消耗
	GetAdAccountReportFieldShowCnt                                    GetAdAccountReportField = "show_cnt"                                          // 展示次数
	GetAdAccountReportFieldCtr                                        GetAdAccountReportField = "ctr"                                               // 点击率
	GetAdAccountReportFieldCpmPlatform                                GetAdAccountReportField = "cpm_platform"                                      // 平均千次展示费用
	GetAdAccountReportFieldClickCnt                                   GetAdAccountReportField = "click_cnt"                                         // 点击次数
	GetAdAccountReportFieldPayOrderCount                              GetAdAccountReportField = "pay_order_count"                                   // 直接成交订单数
	GetAdAccountReportFieldPayOrderAmount                             GetAdAccountReportField = "pay_order_amount"                                  // 直接成交金额
	GetAdAccountReportFieldPrepayAndPayOrderROI                       GetAdAccountReportField = "prepay_and_pay_order_roi"                          // 直接支付roi
	GetAdAccountReportFieldCreateOrderCount                           GetAdAccountReportField = "create_order_count"                                // 直接下单订单数
	GetAdAccountReportFieldCreateOrderAmount                          GetAdAccountReportField = "create_order_amount"                               // 直接下单金额
	GetAdAccountReportFieldCreateOrderROI                             GetAdAccountReportField = "create_order_roi"                                  // 直接下单roi
	GetAdAccountReportFieldPrepayOrderCount                           GetAdAccountReportField = "prepay_order_count"                                // 直接预售订单数
	GetAdAccountReportFieldPrepayOrderAmount                          GetAdAccountReportField = "prepay_order_amount"                               // 直接预售金额
	GetAdAccountReportFieldDyFollow                                   GetAdAccountReportField = "dy_follow"                                         // 新增粉丝数
	GetAdAccountReportFieldTotalPlay                                  GetAdAccountReportField = "total_play"                                        // 播放数
	GetAdAccountReportFieldPlayDuration3s                             GetAdAccountReportField = "play_duration_3s"                                  // 3s播放数
	GetAdAccountReportFieldPlay25FeedBreak                            GetAdAccountReportField = "play_25_feed_break"                                // 25%进度播放数
	GetAdAccountReportFieldPlay50FeedBreak                            GetAdAccountReportField = "play_50_feed_break"                                // 50%进度播放数
	GetAdAccountReportFieldPlay75FeedBreak                            GetAdAccountReportField = "play_75_feed_break"                                // 75%进度播放数
	GetAdAccountReportFieldPlayOver                                   GetAdAccountReportField = "play_over"                                         // 播放完成数
	GetAdAccountReportFieldPlayOverRate                               GetAdAccountReportField = "play_over_rate"                                    // 完播率
	GetAdAccountReportFieldCpcPlatform                                GetAdAccountReportField = "cpc_platform"                                      // 平均点击单价
	GetAdAccountReportFieldDeepConvertCnt                             GetAdAccountReportField = "deep_convert_cnt"                                  // 深度转化次数
	GetAdAccountReportFieldDeepConvertCost                            GetAdAccountReportField = "deep_convert_cost"                                 // 深度转化成本
	GetAdAccountReportFieldDeepConvertRate                            GetAdAccountReportField = "deep_convert_rate"                                 // 深度转化率
	GetAdAccountReportFieldAttributionConvertCnt                      GetAdAccountReportField = "attribution_convert_cnt"                           // 转化数（计费时间）
	GetAdAccountReportFieldAttributionConvertRate                     GetAdAccountReportField = "attribution_convert_rate"                          // 转化率（计费时间）
	GetAdAccountReportFieldAttributionConvertCost                     GetAdAccountReportField = "attribution_convert_cost"                          // 转化成本（计费时间）
	GetAdAccountReportFieldAttributionDeepConvertCnt                  GetAdAccountReportField = "attribution_deep_convert_cnt"                      // 深度转化次数（计费时间）
	GetAdAccountReportFieldAttributionDeepConvertCost                 GetAdAccountReportField = "attribution_deep_convert_cost"                     // 深度转化成本（计费时间）
	GetAdAccountReportFieldAttributionDeepConvertRate                 GetAdAccountReportField = "attribution_deep_convert_rate"                     // 深度转化率（计费时间）
	GetAdAccountReportFieldAllOrderCreateROI7days                     GetAdAccountReportField = "all_order_create_roi_7days"                        // 7日总下单ROI
	GetAdAccountReportFieldAllOrderPayROI7days                        GetAdAccountReportField = "all_order_pay_roi_7days"                           // 7日总支付ROI
	GetAdAccountReportFieldAllOrderPayCount7days                      GetAdAccountReportField = "all_order_pay_count_7days"                         // 7日总成交订单
	GetAdAccountReportFieldAllOrderPayGMV7days                        GetAdAccountReportField = "all_order_pay_gmv_7days"                           // 7日总成交金额
	GetAdAccountReportFieldPayOrderCostPerOrder                       GetAdAccountReportField = "pay_order_cost_per_order"                          // 直接成交客单价
	GetAdAccountReportFieldCreateOrderCouponAmount                    GetAdAccountReportField = "create_order_coupon_amount"                        // 下单智能优惠券金额
	GetAdAccountReportFieldPayOrderCouponAmount                       GetAdAccountReportField = "pay_order_coupon_amount"                           // 成交智能优惠券金额
	GetAdAccountReportFieldIndirectOrderCreateCount7days              GetAdAccountReportField = "indirect_order_create_count_7days"                 // 间接下单订单数
	GetAdAccountReportFieldIndirectOrderCreateGMV7days                GetAdAccountReportField = "indirect_order_create_gmv_7days"                   // 间接下单金额
	GetAdAccountReportFieldIndirectOrderPayCount7days                 GetAdAccountReportField = "indirect_order_pay_count_7days"                    // 间接成交订单数
	GetAdAccountReportFieldIndirectOrderPayGMV7days                   GetAdAccountReportField = "indirect_order_pay_gmv_7days"                      // 间接成交金额
	GetAdAccountReportFieldIndirectOrderPrepayCount7days              GetAdAccountReportField = "indirect_order_prepay_count_7days"                 // 间接预售订单数
	GetAdAccountReportFieldIndirectOrderPrepayGMV7days                GetAdAccountReportField = "indirect_order_prepay_gmv_7days"                   // 间接预售金额
	GetAdAccountReportFieldQianchuanFirstOrderCnt                     GetAdAccountReportField = "qianchuan_first_order_cnt"                         // 店铺首单新客人数
	GetAdAccountReportFieldQianchuanFirstOrderRate                    GetAdAccountReportField = "qianchuan_first_order_rate"                        // 店铺首单新客订单占比
	GetAdAccountReportFieldQianchuanFirstOrderConvertCost             GetAdAccountReportField = "qianchuan_first_order_convert_cost"                // 店铺首单新客转化成本
	GetAdAccountReportFieldQianchuanFirstOrderDirectPayGMV            GetAdAccountReportField = "qianchuan_first_order_direct_pay_gmv"              // 店铺首单新客直接成交金额
	GetAdAccountReportFieldQianchuanFirstOrderDirectPayOrderROI       GetAdAccountReportField = "qianchuan_first_order_direct_pay_order_roi"        // 店铺首单新客直接支付ROI
	GetAdAccountReportFieldQianchuanFirstOrderLTV30                   GetAdAccountReportField = "qianchuan_first_order_ltv30"                       // 店铺首单新客30天累计成交金额
	GetAdAccountReportFieldQianchuanFirstOrderROI30                   GetAdAccountReportField = "qianchuan_first_order_roi30"                       // 店铺首单新客30天累计支付ROI
	GetAdAccountReportFieldQianchuanBrandFirstOrderRate               GetAdAccountReportField = "qianchuan_brand_first_order_rate"                  // 品牌首单新客订单占比
	GetAdAccountReportFieldQianchuanBrandFirstOrderConvertCost        GetAdAccountReportField = "qianchuan_brand_first_order_convert_cost"          // 品牌首单新客转化成本
	GetAdAccountReportFieldQianchuanBrandFirstOrderLTV30              GetAdAccountReportField = "qianchuan_brand_first_order_ltv30"                 // 品牌首单新客30天累计成交金额
	GetAdAccountReportFieldQianchuanBrandFirstOrderDirectPayOrderROI  GetAdAccountReportField = "qianchuan_brand_first_order_direct_pay_order_roi"  // 品牌首单新客直接支付ROI
	GetAdAccountReportFieldQianchuanBrandFirstOrderROI30              GetAdAccountReportField = "qianchuan_brand_first_order_roi30"                 // 品牌首单新客30天累计支付ROI
	GetAdAccountReportFieldQianchuanBrandFirstOrderCnt                GetAdAccountReportField = "qianchuan_brand_first_order_cnt"                   // 品牌首单新客数
	GetAdAccountReportFieldQianchuanBrandFirstOrderDirectPayGMV       GetAdAccountReportField = "qianchuan_brand_first_order_direct_pay_gmv"        // 品牌首单新客直接成交金额
	GetAdAccountReportFieldQianchuanAuthorFirstOrderRate              GetAdAccountReportField = "qianchuan_author_first_order_rate"                 // 抖音号首单新客订单占比
	GetAdAccountReportFieldQianchuanAuthorFirstOrderROI30             GetAdAccountReportField = "qianchuan_author_first_order_roi30"                // 抖音号首单新客30天累计支付ROI
	GetAdAccountReportFieldQianchuanAuthorFirstOrderCnt               GetAdAccountReportField = "qianchuan_author_first_order_cnt"                  // 抖音号首单新客数
	GetAdAccountReportFieldQianchuanAuthorFirstOrderConvertCost       GetAdAccountReportField = "qianchuan_author_first_order_convert_cost"         // 抖音号首单新客转化成本
	GetAdAccountReportFieldQianchuanAuthorFirstOrderDirectPayGMV      GetAdAccountReportField = "qianchuan_author_first_order_direct_pay_gmv"       // 抖音号首单新客直接成交金额
	GetAdAccountReportFieldQianchuanAuthorFirstOrderDirectPayOrderROI GetAdAccountReportField = "qianchuan_author_first_order_direct_pay_order_roi" // 抖音号首单新客直接支付ROI
	GetAdAccountReportFieldQianchuanAuthorFirstOrderLTV30             GetAdAccountReportField = "qianchuan_author_first_order_ltv30"                // 抖音号首单新客30天累计成交金额
	GetAdAccountReportFieldLiveOrderSettleAmount14D                   GetAdAccountReportField = "live_order_settle_amount_14d"                      // 14天结算金额
	GetAdAccountReportFieldLiveOrderSettleCount14D                    GetAdAccountReportField = "live_order_settle_count_14d"                       // 14天结算订单数
	GetAdAccountReportFieldLiveOrderSettleCountRate14D                GetAdAccountReportField = "live_order_settle_count_rate_14d"                  // 14天结算订单率
	GetAdAccountReportFieldAdLiveOrderSettleROI14D                    GetAdAccountReportField = "ad_live_order_settle_roi_14d"                      // 14天结算ROI
	GetAdAccountReportFieldAdLiveOrderSettleCost14D                   GetAdAccountReportField = "ad_live_order_settle_cost_14d"                     // 14天结算成本
	GetAdAccountReportFieldUnfinishedEstimateOrderGMV                 GetAdAccountReportField = "unfinished_estimate_order_gmv"                     // 未完结直接预售订单预估金额
	GetAdAccountReportFieldIndirectOrderUnfinishedEstimateGMV7Days    GetAdAccountReportField = "indirect_order_unfinished_estimate_gmv_7days"      // 未完结间接预售订单预估金额
)
