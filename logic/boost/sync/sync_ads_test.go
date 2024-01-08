package sync

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"testing"
)

func init() {
	dal.Init()
	redis_dal.Init()
}

func TestSyncAds(t *testing.T) {
	type args struct {
		ctx          context.Context
		advertiserID int64
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
				advertiserID: 1703886601680909,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SyncAds(tt.args.ctx, tt.args.advertiserID); (err != nil) != tt.wantErr {
				t.Errorf("SyncAds() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
