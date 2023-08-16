package wallet

import (
	"context"
	"testing"
)

func TestGetWallet(t *testing.T) {
	type args struct {
		ctx  context.Context
		adID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:  context.Background(),
				adID: 1748031128935424, //我们的号:1748031128935424
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetWallet(tt.args.ctx, tt.args.adID); (err != nil) != tt.wantErr {
				t.Errorf("GetWallet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
