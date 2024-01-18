package ad_group

import (
	"context"
	"reflect"
	"testing"
)

func TestUpdateAdGroup(t *testing.T) {
	type args struct {
		ctx context.Context
		req *UpdateAdGroupReq
	}
	tests := []struct {
		name    string
		args    args
		want    *UpdateAdGroupRespData
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				req: &UpdateAdGroupReq{
					AdvertiserID: 1784698853978186,
					CampaignID:   1788144214349987,
					BudgetMode:   "",
					Budget:       "",
					CampaignName: "测试",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateAdGroup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAdGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}
