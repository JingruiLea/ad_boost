package ad_group

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

func TestCreateAdGroup(t *testing.T) {
	type args struct {
		ctx context.Context
		req *CreateAdGroupReq
	}
	tests := []struct {
		name    string
		args    args
		want    *CreateAdGroupRespData
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				req: &CreateAdGroupReq{
					AdvertiserID:   1784698853978186,
					CampaignName:   "千川机器人",
					MarketingGoal:  ttypes.MarketingGoalLivePromGoods,
					MarketingScene: ttypes.MarketingSceneFeed,
					BudgetMode:     ttypes.BudgetModeInfinite,
					Budget:         0,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAdGroup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAdGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAdGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}
