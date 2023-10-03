package ad

import (
	"context"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"reflect"
	"testing"
)

func TestGetAdList(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetAdListReq
	}
	tests := []struct {
		name    string
		args    args
		want    *GetAdListRespData
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				req: &GetAdListReq{
					AdvertiserId:     1748031128935424,
					RequestAwemeInfo: 0,
					Filtering: &Filter{
						AdName:        "推直播间",
						MarketingGoal: ttypes.MarketingGoalLivePromGoods,
					},
					Page:     0,
					PageSize: 0,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAdList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdList() got = %v, want %v", got, tt.want)
			}
			adIDs := make([]int64, 0, len(got.List))
			for _, ad := range got.List {
				adIDs = append(adIDs, ad.AdID)
			}
			t.Logf("adIDs: %v", utils.GetJsonStr(adIDs))
		})
	}
}
