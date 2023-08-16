package ad_group

import (
	"context"
	"testing"
)

func TestGetAdGroupList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetAdGroupListReq
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
				req: &GetAdGroupListReq{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GetAdGroupList(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("GetAdGroupList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
