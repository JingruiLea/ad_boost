package account

import (
	"context"
	"reflect"
	"testing"
)

func TestGetAccountType(t *testing.T) {
	type args struct {
		ctx        context.Context
		accountIDs []int64
	}
	tests := []struct {
		name    string
		args    args
		want    *GetAccountTypeRespData
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:        context.Background(),
				accountIDs: []int64{1703886601680909},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAccountType(tt.args.ctx, tt.args.accountIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountType() got = %v, want %v", got, tt.want)
			}
		})
	}
}
