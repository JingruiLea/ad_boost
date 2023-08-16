package live_report

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
	"time"
)

func GetReport(ctx context.Context, req *GetReportReq) error {
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().Get(ctx, "https://api.oceanengine.com/open_api/v1.0/qianchuan/report/live/get/", httpclient.CommonHeader, &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetReport httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetReport respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}

type GetReportReq struct {
	AdvertiserID   int64              `json:"advertiser_id"`             // 广告主id
	AwemeID        int64              `json:"aweme_id"`                  // 抖音号ID
	StartTime      string             `json:"start_time"`                // 开始时间
	EndTime        string             `json:"end_time"`                  // 结束时间
	Fields         []RoomMetricsField `json:"fields"`                    // 需要查询的消耗指标
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
		EndTime:      endTime.Format("2006-01-02 15:04:05"),
		Fields:       fields,
	}
}

type RoomMetricsField string

// 定义所有字段的常量
const (
	TatCost                       RoomMetricsField = "tat_cost"
	CpmPlatform                   RoomMetricsField = "cpm_platform"
	ClickCnt                      RoomMetricsField = "click_cnt"
	Ctr                           RoomMetricsField = "ctr"
	TotalLivePayOrderGpm          RoomMetricsField = "total_live_pay_order_gpm"
	LubanLivePayOrderGpm          RoomMetricsField = "luban_live_pay_order_gpm"
	CpcPlatform                   RoomMetricsField = "cpc_platform"
	ConvertCnt                    RoomMetricsField = "convert_cnt"
	ConvertRate                   RoomMetricsField = "convert_rate"
	CpaPlatform                   RoomMetricsField = "cpa_platform"
	LivePayOrderGmvAlias          RoomMetricsField = "live_pay_order_gmv_alias"
	LubanLivePayOrderGmv          RoomMetricsField = "luban_live_pay_order_gmv"
	LivePayOrderGmvRoi            RoomMetricsField = "live_pay_order_gmv_roi"
	AdLivePrepayAndPayOrderGmvRoi RoomMetricsField = "ad_live_prepay_and_pay_order_gmv_roi"
	LiveCreateOrderCountAlias     RoomMetricsField = "live_create_order_count_alias"
	LiveCreateOrderRate           RoomMetricsField = "live_create_order_rate"
	LubanLiveOrderCount           RoomMetricsField = "luban_live_order_count"
	AdLiveCreateOrderRate         RoomMetricsField = "ad_live_create_order_rate"
	LivePayOrderCountAlias        RoomMetricsField = "live_pay_order_count_alias"
	LivePayOrderRate              RoomMetricsField = "live_pay_order_rate"
	LubanLivePayOrderCount        RoomMetricsField = "luban_live_pay_order_count"
	AdLivePayOrderRate            RoomMetricsField = "ad_live_pay_order_rate"
	LivePayOrderGmvAvg            RoomMetricsField = "live_pay_order_gmv_avg"
	AdLivePayOrderGmvAvg          RoomMetricsField = "ad_live_pay_order_gmv_avg"
	LubanLivePrepayOrderCount     RoomMetricsField = "luban_live_prepay_order_count"
	LubanLivePrepayOrderGmv       RoomMetricsField = "luban_live_prepay_order_gmv"
	LivePrepayOrderCountAlias     RoomMetricsField = "live_prepay_order_count_alias"
	LivePrepayOrderGmvAlias       RoomMetricsField = "live_prepay_order_gmv_alias"
	LiveOrderPayCouponAmount      RoomMetricsField = "live_order_pay_coupon_amount"
	TotalLiveWatchCnt             RoomMetricsField = "total_live_watch_cnt"
	TotalLiveFollowCnt            RoomMetricsField = "total_live_follow_cnt"
	LiveWatchOneMinuteCount       RoomMetricsField = "live_watch_one_minute_count"
	TotalLiveFansClubJoinCnt      RoomMetricsField = "total_live_fans_club_join_cnt"
	LiveClickCartCountAlias       RoomMetricsField = "live_click_cart_count_alias"
	LiveClickProductCountAlias    RoomMetricsField = "live_click_product_count_alias"
	TotalLiveCommentCnt           RoomMetricsField = "total_live_comment_cnt"
	TotalLiveShareCnt             RoomMetricsField = "total_live_share_cnt"
	TotalLiveGiftCnt              RoomMetricsField = "total_live_gift_cnt"
	TotalLiveGiftAmount           RoomMetricsField = "total_live_gift_amount"
)

// All 返回一个包含所有字段的字符串切片
func (*RoomMetricsField) All() []RoomMetricsField {
	return []RoomMetricsField{
		TatCost,
		CpmPlatform,
		ClickCnt,
		Ctr,
		TotalLivePayOrderGpm,
		LubanLivePayOrderGpm,
		CpcPlatform,
		ConvertCnt,
		ConvertRate,
		CpaPlatform,
		LivePayOrderGmvAlias,
		LubanLivePayOrderGmv,
		LivePayOrderGmvRoi,
		AdLivePrepayAndPayOrderGmvRoi,
		LiveCreateOrderCountAlias,
		LiveCreateOrderRate,
		LubanLiveOrderCount,
		AdLiveCreateOrderRate,
		LivePayOrderCountAlias,
		LivePayOrderRate,
		LubanLivePayOrderCount,
		AdLivePayOrderRate,
		LivePayOrderGmvAvg,
		AdLivePayOrderGmvAvg,
		LubanLivePrepayOrderCount,
		LubanLivePrepayOrderGmv,
		LivePrepayOrderCountAlias,
		LivePrepayOrderGmvAlias,
		LiveOrderPayCouponAmount,
		TotalLiveWatchCnt,
		TotalLiveFollowCnt,
		LiveWatchOneMinuteCount,
		TotalLiveFansClubJoinCnt,
		LiveClickCartCountAlias,
		LiveClickProductCountAlias,
		TotalLiveCommentCnt,
		TotalLiveShareCnt,
		TotalLiveGiftCnt,
		TotalLiveGiftAmount,
	}
}
