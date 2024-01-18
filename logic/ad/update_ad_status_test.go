package ad

import (
	"context"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"testing"
)

func TestUpdateAdStatus(t *testing.T) {
	type args struct {
		ctx context.Context
		req *UpdateAdStatusReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				req: &UpdateAdStatusReq{
					AdvertiserID: 1784698853978186,
					AdIDs:        []int64{},
					OptStatus:    ttypes.OptStatusDisable,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateAdStatus(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
