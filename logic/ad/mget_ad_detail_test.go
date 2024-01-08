package ad

import (
	"context"
	"reflect"
	"testing"
)

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
				advertiserID: 1703886601680909,
				adID:         0,
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
