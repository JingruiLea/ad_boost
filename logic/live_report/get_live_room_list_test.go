package live_report

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"testing"
)

func init() {
	redis_dal.Init()
	dal.Init()
}

func TestGetLiveRoomList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetLiveRoomListReq
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
				req: &GetLiveRoomListReq{
					AdvertiserID: 1703886601680909,
					AwemeID:      2691211639665967,
					DateTime:     "2024-01-07",
					RoomStatus:   RoomStatusAll,
					AdStatus:     AdStatusAll,
					Fields:       RoomMetricsFieldStatCost.All(),
					Page:         1,
					PageSize:     10,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GetLiveRoomList(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("GetLiveRoomList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
