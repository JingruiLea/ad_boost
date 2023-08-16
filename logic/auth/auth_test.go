package auth

import (
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"golang.org/x/net/context"
	"testing"
)

func TestAuth(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Auth(tt.args.ctx)
		})
	}
}

//{"code":0,"data":{"access_token":"c222849a000dbda9ff4e73d17d45180ecfd52579","expires_in":86399,"refresh_token":"35df4b0fa733e0d449e04749bec0620d1c9a11d8","refresh_token_expires_in":2591999},"message":"OK","request_id":"202308112004345CDF966E6DCA9FB1FD1E"}

func TestRefreshToken(t *testing.T) {
	redis_dal.Init()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantAt  string
		wantRt  string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
			},
			wantAt:  "",
			wantRt:  "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAt, gotRt, err := RefreshToken(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAt != tt.wantAt {
				t.Errorf("RefreshToken() gotAt = %v, want %v", gotAt, tt.wantAt)
			}
			if gotRt != tt.wantRt {
				t.Errorf("RefreshToken() gotRt = %v, want %v", gotRt, tt.wantRt)
			}
		})
	}
}
