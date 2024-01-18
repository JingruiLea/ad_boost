package ad

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"reflect"
	"testing"
)

func init() {
	dal.Init()
	redis_dal.Init()
}

func TestGetAdDetail(t *testing.T) {
	type args struct {
		ctx          context.Context
		advertiserID int64
		adID         int64
	}
	tests := []struct {
		name    string
		args    args
		want    *GetAdDetailRespData
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.Background(),
				advertiserID: 1786309327331339,
				adID:         1787688640320515,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAdDetail(tt.args.ctx, tt.args.advertiserID, tt.args.adID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
