package sync

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/account_dal"
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
				advertiserID: 1784698853978186,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ctx := tt.args.ctx
		accounts, err := account_dal.MGetAllAccount(ctx)
		if err != nil {
			t.Errorf("SyncAds() error = %v", err)
			return
		}
		t.Run(tt.name, func(t *testing.T) {
			for _, account := range accounts {
				if err := SyncAds(tt.args.ctx, account.AdvertiserID); (err != nil) != tt.wantErr {
					t.Errorf("SyncAds() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
