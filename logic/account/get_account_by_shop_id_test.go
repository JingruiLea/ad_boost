package account

import (
	"context"
	"reflect"
	"testing"
)

func TestGetAdAccountByShopID(t *testing.T) {
	type args struct {
		ctx    context.Context
		shopID int64
	}
	tests := []struct {
		name     string
		args     args
		wantData []int64
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				ctx:    context.Background(),
				shopID: 19642944,
			},
			wantData: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := GetAdAccountByShopID(tt.args.ctx, tt.args.shopID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdAccountByShopID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("GetAdAccountByShopID() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
