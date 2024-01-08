package ad_report

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"reflect"
	"testing"
)

func init() {
	redis_dal.Init()
}

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

func TestMGetCommonAdReport(t *testing.T) {
	type args struct {
		ctx          context.Context
		advertiserID int64
		adIDs        []int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*AdReport
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.Background(),
				advertiserID: 1703886601680909,
				adIDs: []int64{
					1787256711414809,
					1787256710989923,
					1787256710709305,
					1787256710342723,
					1787254871200884,
					1787254275657851,
					1787153238160425,
					1787153189334020,
					1787153090339924,
					1787153053722660,
					1787153035385947,
					1787152993963155,
					1787152967562404,
					1787152945139764,
					1787150360981689,
					1787150313979987,
					1787150273592435,
					1787150082496524,
					1787149678188596,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MGetCommonAdDailyReport(tt.args.ctx, tt.args.advertiserID, tt.args.adIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("MGetCommonAdReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MGetCommonAdReport() got = %v, want %v", got, tt.want)
			}
		})
	}
}
