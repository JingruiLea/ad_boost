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
				accessToken:  "a3921c59e232fef3612a17390c871d54f89e365a",
				refreshToken: "4316e1bfc5e7d99f52880ad21393e212548ea0e0",
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
