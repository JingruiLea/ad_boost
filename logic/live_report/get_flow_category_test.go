package live_report

import (
	"context"
	"reflect"
	"testing"
)

func TestGetFlowCategory(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetFlowCategoryReq
	}
	tests := []struct {
		name    string
		args    args
		want    *GetFlowCategoryRespData
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				req: &GetFlowCategoryReq{
					AdvertiserID: 1703886601680909,
					StartTime:    nil,
					EndTime:      nil,
					Fields:       nil,
					Filtering:    Filtering{},
					OrderField:   nil,
					OrderType:    nil,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFlowCategory(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFlowCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFlowCategory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
