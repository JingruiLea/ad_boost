package live_report

import (
	"context"
	"testing"
)

func TestGetRoomDetail(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetRoomDetailReq
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
				req: &GetRoomDetailReq{
					AdvertiserID: 1748031128935424,
					RoomID:       7267915748029238077,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetRoomDetail(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("GetRoomDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
