package sync

import (
	"context"
	"github.com/JingruiLea/ad_boost/logic/ad_group"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"testing"
)

func TestSyncAdGroup(t *testing.T) {
	type args struct {
		ctx          context.Context
		advertiserID int64
		filter       *ad_group.Filter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.Background(),
				advertiserID: 1784698853978186,
				filter: &ad_group.Filter{
					MarketingGoal: ttypes.MarketingGoalLivePromGoods,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SyncAdGroup(tt.args.ctx, tt.args.advertiserID, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("SyncAdGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
