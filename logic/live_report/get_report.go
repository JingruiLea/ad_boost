package live_report

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
	"time"
)

func GetReport(ctx context.Context, req *GetReportReq) (*GetReportRespData, error) {
	var resp GetReportRespData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://api.oceanengine.com/open_api/v1.0/qianchuan/report/live/get/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetReport httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	fmt.Printf("GetReport respMap: %s", utils.GetJsonStr(resp))
	return &resp, nil
}

type GetReportReq struct {
	AdvertiserID   int64              `json:"advertiser_id"`             // 广告主id
	AwemeID        int64              `json:"aweme_id"`                  // 抖音号ID
	StartTime      string             `json:"start_time"`                // 开始时间
	EndTime        string             `json:"end_time"`                  // 结束时间
	Fields         []RoomMetricsField `json:"fields,omitempty"`          // 需要查询的消耗指标
	StatsAuthority StatsAuthority     `json:"stats_authority,omitempty"` // 广告账户维度
}

type StatsAuthority string

const (
	StatsAuthorityQualification = "QUALIFICATION"
	StatsAuthorityCurrent       = "CURRENT"
)

func NewGetReportReq(adID, awemeID int64, startTime, endTime time.Time, fields []RoomMetricsField) *GetReportReq {
	return &GetReportReq{
		AdvertiserID: adID,
		AwemeID:      awemeID,
		StartTime:    startTime.Format("2006-01-02 15:04:05"),
		EndTime:      endTime.AddDate(0, 0, -1).Format("2006-01-02 15:04:05"),
		Fields:       fields,
	}
}

type RoomMetricsField string

// 定义所有字段的常量
const (
	RoomMetricsFieldStatCost                      RoomMetricsField = "stat_cost"
	RoomMetricsFieldCpmPlatform                   RoomMetricsField = "cpm_platform"
	RoomMetricsFieldClickCnt                      RoomMetricsField = "click_cnt"
	RoomMetricsFieldCtr                           RoomMetricsField = "ctr"
	RoomMetricsFieldTotalLivePayOrderGpm          RoomMetricsField = "total_live_pay_order_gpm"
	RoomMetricsFieldLubanLivePayOrderGpm          RoomMetricsField = "luban_live_pay_order_gpm"
	RoomMetricsFieldCpcPlatform                   RoomMetricsField = "cpc_platform"
	RoomMetricsFieldConvertCnt                    RoomMetricsField = "convert_cnt"
	RoomMetricsFieldConvertRate                   RoomMetricsField = "convert_rate"
	RoomMetricsFieldCpaPlatform                   RoomMetricsField = "cpa_platform"
	RoomMetricsFieldLivePayOrderGmvAlias          RoomMetricsField = "live_pay_order_gmv_alias"
	RoomMetricsFieldLubanLivePayOrderGmv          RoomMetricsField = "luban_live_pay_order_gmv"
	RoomMetricsFieldLivePayOrderGmvRoi            RoomMetricsField = "live_pay_order_gmv_roi"
	RoomMetricsFieldAdLivePrepayAndPayOrderGmvRoi RoomMetricsField = "ad_live_prepay_and_pay_order_gmv_roi"
	RoomMetricsFieldLiveCreateOrderCountAlias     RoomMetricsField = "live_create_order_count_alias"
	RoomMetricsFieldLiveCreateOrderRate           RoomMetricsField = "live_create_order_rate"
	RoomMetricsFieldLubanLiveOrderCount           RoomMetricsField = "luban_live_order_count"
	RoomMetricsFieldAdLiveCreateOrderRate         RoomMetricsField = "ad_live_create_order_rate"
	RoomMetricsFieldLivePayOrderCountAlias        RoomMetricsField = "live_pay_order_count_alias"
	RoomMetricsFieldLivePayOrderRate              RoomMetricsField = "live_pay_order_rate"
	RoomMetricsFieldLubanLivePayOrderCount        RoomMetricsField = "luban_live_pay_order_count"
	RoomMetricsFieldAdLivePayOrderRate            RoomMetricsField = "ad_live_pay_order_rate"
	RoomMetricsFieldLivePayOrderGmvAvg            RoomMetricsField = "live_pay_order_gmv_avg"
	RoomMetricsFieldAdLivePayOrderGmvAvg          RoomMetricsField = "ad_live_pay_order_gmv_avg"
	RoomMetricsFieldLubanLivePrepayOrderCount     RoomMetricsField = "luban_live_prepay_order_count"
	RoomMetricsFieldLubanLivePrepayOrderGmv       RoomMetricsField = "luban_live_prepay_order_gmv"
	RoomMetricsFieldLivePrepayOrderCountAlias     RoomMetricsField = "live_prepay_order_count_alias"
	RoomMetricsFieldLivePrepayOrderGmvAlias       RoomMetricsField = "live_prepay_order_gmv_alias"
	RoomMetricsFieldLiveOrderPayCouponAmount      RoomMetricsField = "live_order_pay_coupon_amount"
	RoomMetricsFieldTotalLiveWatchCnt             RoomMetricsField = "total_live_watch_cnt"
	RoomMetricsFieldTotalLiveFollowCnt            RoomMetricsField = "total_live_follow_cnt"
	RoomMetricsFieldLiveWatchOneMinuteCount       RoomMetricsField = "live_watch_one_minute_count"
	RoomMetricsFieldTotalLiveFansClubJoinCnt      RoomMetricsField = "total_live_fans_club_join_cnt"
	RoomMetricsFieldLiveClickCartCountAlias       RoomMetricsField = "live_click_cart_count_alias"
	RoomMetricsFieldLiveClickProductCountAlias    RoomMetricsField = "live_click_product_count_alias"
	RoomMetricsFieldTotalLiveCommentCnt           RoomMetricsField = "total_live_comment_cnt"
	RoomMetricsFieldTotalLiveShareCnt             RoomMetricsField = "total_live_share_cnt"
	RoomMetricsFieldTotalLiveGiftCnt              RoomMetricsField = "total_live_gift_cnt"
	RoomMetricsFieldTotalLiveGiftAmount           RoomMetricsField = "total_live_gift_amount"
)

// All 返回一个包含所有字段的字符串切片
func (RoomMetricsField) All() []RoomMetricsField {
	return []RoomMetricsField{
		RoomMetricsFieldStatCost,
		RoomMetricsFieldCpmPlatform,
		RoomMetricsFieldClickCnt,
		RoomMetricsFieldCtr,
		RoomMetricsFieldTotalLivePayOrderGpm,
		RoomMetricsFieldLubanLivePayOrderGpm,
		RoomMetricsFieldCpcPlatform,
		RoomMetricsFieldConvertCnt,
		RoomMetricsFieldConvertRate,
		RoomMetricsFieldCpaPlatform,
		RoomMetricsFieldLivePayOrderGmvAlias,
		RoomMetricsFieldLubanLivePayOrderGmv,
		RoomMetricsFieldLivePayOrderGmvRoi,
		RoomMetricsFieldAdLivePrepayAndPayOrderGmvRoi,
		RoomMetricsFieldLiveCreateOrderCountAlias,
		RoomMetricsFieldLiveCreateOrderRate,
		RoomMetricsFieldLubanLiveOrderCount,
		RoomMetricsFieldAdLiveCreateOrderRate,
		RoomMetricsFieldLivePayOrderCountAlias,
		RoomMetricsFieldLivePayOrderRate,
		RoomMetricsFieldLubanLivePayOrderCount,
		RoomMetricsFieldAdLivePayOrderRate,
		RoomMetricsFieldLivePayOrderGmvAvg,
		RoomMetricsFieldAdLivePayOrderGmvAvg,
		RoomMetricsFieldLubanLivePrepayOrderCount,
		RoomMetricsFieldLubanLivePrepayOrderGmv,
		RoomMetricsFieldLivePrepayOrderCountAlias,
		RoomMetricsFieldLivePrepayOrderGmvAlias,
		RoomMetricsFieldLiveOrderPayCouponAmount,
		RoomMetricsFieldTotalLiveWatchCnt,
		RoomMetricsFieldTotalLiveFollowCnt,
		RoomMetricsFieldLiveWatchOneMinuteCount,
		RoomMetricsFieldTotalLiveFansClubJoinCnt,
		RoomMetricsFieldLiveClickCartCountAlias,
		RoomMetricsFieldLiveClickProductCountAlias,
		RoomMetricsFieldTotalLiveCommentCnt,
		RoomMetricsFieldTotalLiveShareCnt,
		RoomMetricsFieldTotalLiveGiftCnt,
		RoomMetricsFieldTotalLiveGiftAmount,
	}
}

type GetReportRespData struct {
	AdLiveCreateOrderRate         float64 `json:"ad_live_create_order_rate"`
	AdLivePayOrderGmvAvg          float64 `json:"ad_live_pay_order_gmv_avg"`
	AdLivePayOrderRate            float64 `json:"ad_live_pay_order_rate"`
	AdLivePrepayAndPayOrderGmvRoi float64 `json:"ad_live_prepay_and_pay_order_gmv_roi"`
	ClickCnt                      int     `json:"click_cnt"`
	ConvertCnt                    int     `json:"convert_cnt"`
	ConvertRate                   float64 `json:"convert_rate"`
	CpaPlatform                   float64 `json:"cpa_platform"`
	CpcPlatform                   float64 `json:"cpc_platform"`
	CpmPlatform                   float64 `json:"cpm_platform"`
	Ctr                           float64 `json:"ctr"`
	LiveClickCartCountAlias       int     `json:"live_click_cart_count_alias"`
	LiveClickProductCountAlias    int     `json:"live_click_product_count_alias"`
	LiveCreateOrderCountAlias     int     `json:"live_create_order_count_alias"`
	LiveCreateOrderRate           float64 `json:"live_create_order_rate"`
	LiveOrderPayCouponAmount      float64 `json:"live_order_pay_coupon_amount"`
	LivePayOrderCountAlias        int     `json:"live_pay_order_count_alias"`
	LivePayOrderGmvAlias          int64   `json:"live_pay_order_gmv_alias"`
	LivePayOrderGmvAvg            float64 `json:"live_pay_order_gmv_avg"`
	LivePayOrderGmvRoi            float64 `json:"live_pay_order_gmv_roi"`
	LivePayOrderRate              float64 `json:"live_pay_order_rate"`
	LivePrepayOrderCountAlias     int     `json:"live_prepay_order_count_alias"`
	LivePrepayOrderGmvAlias       int     `json:"live_prepay_order_gmv_alias"`
	LiveWatchOneMinuteCount       int     `json:"live_watch_one_minute_count"`
	LubanLiveOrderCount           int     `json:"luban_live_order_count"`
	LubanLivePayOrderCount        int     `json:"luban_live_pay_order_count"`
	LubanLivePayOrderGmv          float64 `json:"luban_live_pay_order_gmv"`
	LubanLivePayOrderGpm          float64 `json:"luban_live_pay_order_gpm"`
	LubanLivePrepayOrderCount     int     `json:"luban_live_prepay_order_count"`
	LubanLivePrepayOrderGmv       float64 `json:"luban_live_prepay_order_gmv"`
	QualificationStatCost         int     `json:"qualification_stat_cost"`
	StatCost                      float64 `json:"stat_cost"`
	TotalLiveCommentCnt           int     `json:"total_live_comment_cnt"`
	TotalLiveFansClubJoinCnt      int     `json:"total_live_fans_club_join_cnt"`
	TotalLiveFollowCnt            int     `json:"total_live_follow_cnt"`
	TotalLiveGiftAmount           float64 `json:"total_live_gift_amount"`
	TotalLiveGiftCnt              float64 `json:"total_live_gift_cnt"`
	TotalLivePayOrderGpm          float64 `json:"total_live_pay_order_gpm"`
	TotalLiveShareCnt             int     `json:"total_live_share_cnt"`
	TotalLiveWatchCnt             int     `json:"total_live_watch_cnt"`
}
