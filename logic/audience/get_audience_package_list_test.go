package audience

import (
	"context"
	"testing"
)

func TestGetAudiencePackageList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetAudiencePackageListReq
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
				req: &GetAudiencePackageListReq{
					AdvertiserID: 1784698853978186,
					Filtering:    nil,
					Page:         1,
					PageSize:     100,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GetAudiencePackageList(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("GetAudiencePackageList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
