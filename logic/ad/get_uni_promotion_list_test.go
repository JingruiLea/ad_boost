package ad

import (
	"context"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/jinzhu/now"
	"reflect"
	"testing"
)

func TestGetUniPromotionList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetUniPromotionListReq
	}
	tests := []struct {
		name    string
		args    args
		want    *GetUniPromotionListRespData
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				req: &GetUniPromotionListReq{
					AdvertiserID:  1703886601680909,
					StartTime:     now.BeginningOfDay().Format("2006-01-02 15:04:05"),
					EndTime:       now.EndOfDay().Format("2006-01-02 15:04:05"),
					MarketingGoal: ttypes.MarketingGoalLivePromGoods,
					Fields:        UniPromotionStatFieldStatCost.All(),
					OrderType:     nil,
					OrderField:    nil,
					Page:          nil,
					PageSize:      nil,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUniPromotionList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUniPromotionList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUniPromotionList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
