package ad_report

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"reflect"
	"testing"
)

func init() {
	dal.Init()
	redis_dal.Init()
}

func TestGetAdAccountReport(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetAdAccountReportReq
	}
	tests := []struct {
		name    string
		args    args
		want    *AdvertiserReportRespData
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				req: &GetAdAccountReportReq{
					AdvertiserID:    1772449904394254,
					StartDate:       "2024-01-01",
					EndDate:         "2024-01-07",
					TimeGranularity: nil,
					Fields:          GetAdAccountReportFieldAdLiveOrderSettleCost14D.Common(),
					Filtering: GetAdAccountReportFilter{
						MarketingGoal:  ttypes.MarketingGoalLiveAll,
						OrderPlatform:  nil,
						MarketingScene: nil,
						CampaignScene:  nil,
						SmartBidType:   nil,
						Status:         nil,
						AwemeIDs:       nil,
					},
					OrderField: nil,
					OrderType:  nil,
					Page:       nil,
					PageSize:   nil,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAdAccountReport(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdAccountReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdAccountReport() got = %v, want %v", got, tt.want)
			}
		})
	}
}
