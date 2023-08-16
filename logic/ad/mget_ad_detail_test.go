package ad

import (
	"context"
	"testing"
)

func TestGetAdDetail(t *testing.T) {
	type args struct {
		ctx          context.Context
		advertiserID int64
		adID         int64
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
				adID:         1773119076620308,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetAdDetail(tt.args.ctx, tt.args.advertiserID, tt.args.adID); (err != nil) != tt.wantErr {
				t.Errorf("GetAdDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
