package sync

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"testing"
)

func init() {
	dal.Init()
}

func TestSync(t *testing.T) {
	type args struct {
		ctx          context.Context
		accessToken  string
		refreshToken string
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
				accessToken:  "b7f1815041aa3835b1e8dcc4ede24e7a33cd103e",
				refreshToken: "ff0346703019c70ac8d9232499278a70ad747b39",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SyncAccount(tt.args.ctx, tt.args.accessToken, tt.args.refreshToken); (err != nil) != tt.wantErr {
				t.Errorf("SyncAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
