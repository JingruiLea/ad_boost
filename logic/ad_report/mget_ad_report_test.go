package ad_report

import (
	"context"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"testing"
)

func TestMGetAdReport(t *testing.T) {
	type args struct {
		ctx context.Context
		req *MGetAdReportReq
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
				req: &MGetAdReportReq{
					AdvertiserID: 1748031128935424,
					StartDate:    "2023-08-18",
					EndDate:      "2023-08-18",
					Fields:       MGetAdReportFieldAllOrderCreateRoi7Days.Common(),
					Filtering: &MGetAdReportFiltering{
						MarketingGoal: ttypes.MarketingGoalLivePromGoods,
					},
					Page:     1,
					PageSize: 10,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := MGetAdReport(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("MGetAdReport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
