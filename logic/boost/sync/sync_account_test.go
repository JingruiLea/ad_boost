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
				accessToken:  "c2cc09e50e011f5aefecd5bac1dcdf71a7172315",
				refreshToken: "28e1f7407df9775e74b696d10274bbc71dcad7cd",
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
