package live_report

import (
	"context"
	"testing"
	"time"
)

func TestGetReport(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetReportReq
	}
	fields := []RoomMetricsField{
		"tat_cost",
		"cpm_platform",
		"click_cnt",
		"ctr",
		"total_live_pay_order_gpm",
		"luban_live_pay_order_gpm",
		"cpc_platform",
		"convert_cnt",
		"convert_rate",
		"cpa_platform",
		"live_pay_order_gmv_alias",
		"luban_live_pay_order_gmv",
		"live_pay_order_gmv_roi",
		"ad_live_prepay_and_pay_order_gmv_roi",
		"live_create_order_count_alias",
		"live_create_order_rate",
		"luban_live_order_count",
		"ad_live_create_order_rate",
		"live_pay_order_count_alias",
		"live_pay_order_rate",
		"luban_live_pay_order_count",
		"ad_live_pay_order_rate",
		"live_pay_order_gmv_avg",
		"ad_live_pay_order_gmv_avg",
		"luban_live_prepay_order_count",
		"luban_live_prepay_order_gmv",
		"live_prepay_order_count_alias",
		"live_prepay_order_gmv_alias",
		"live_order_pay_coupon_amount",
		"total_live_watch_cnt",
		"total_live_follow_cnt",
		"live_watch_one_minute_count",
		"total_live_fans_club_join_cnt",
		"live_click_cart_count_alias",
		"live_click_product_count_alias",
		"total_live_comment_cnt",
		"total_live_share_cnt",
		"total_live_gift_cnt",
		"total_live_gift_amount",
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				req: NewGetReportReq(1748031128935424, 2893532936291624, time.Now().Add(time.Hour*48*-1), time.Now(), fields),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetReport(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("GetReport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
