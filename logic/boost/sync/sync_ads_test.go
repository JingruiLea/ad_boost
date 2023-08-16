package sync

import (
	"context"
	"testing"
)

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
				advertiserID: 1748031128935424,
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
