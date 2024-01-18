package audience

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"testing"
)

func init() {
	dal.Init()
	redis_dal.Init()
}

func TestGetDMPPackageList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetDMPPackageListReq
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
				req: &GetDMPPackageListReq{
					AdvertiserId:        1784698853978186,
					RetargetingTagsType: ttypes.RetargetingTagsTypeAll,
					Offset:              0,
					Limit:               100,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetDMPPackageList(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("GetDMPPackageList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
