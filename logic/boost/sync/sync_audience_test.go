package sync

import (
	"context"
	"github.com/JingruiLea/ad_boost/logic/ad_group"
	"testing"
)

func TestSyncAudiencePackage(t *testing.T) {
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
				advertiserID: 1748031128935424,
				filter:       nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SyncAudiencePackage(tt.args.ctx, tt.args.advertiserID, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("SyncAudiencePackage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
