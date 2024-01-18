package tools

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"reflect"
	"testing"
)

func init() {
	dal.Init()
	redis_dal.Init()
}

func TestGetActionCategory(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetActionCategoryReq
	}
	tests := []struct {
		name    string
		args    args
		want    *GetActionCategoryData
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				req: &GetActionCategoryReq{
					AdvertiserID: 1784698853978186,
					ActionScene: []ttypes.ActionScene{
						ttypes.ActionSceneECommerce,
					},
					ActionDays: ttypes.ActionDays180,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetActionCategory(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActionCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActionCategory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
