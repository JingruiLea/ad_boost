package account

import (
	"golang.org/x/net/context"
	"reflect"
	"testing"
)

func TestGetShopAccount(t *testing.T) {
	type args struct {
		ctx context.Context
		at  string
	}
	tests := []struct {
		name         string
		args         args
		wantAccounts []*Account
		wantErr      bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				at:  "c2cc09e50e011f5aefecd5bac1dcdf71a7172315",
			},
			wantAccounts: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAccounts, err := GetShopAccount(tt.args.ctx, tt.args.at)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShopAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAccounts, tt.wantAccounts) {
				t.Errorf("GetShopAccount() gotAccounts = %v, want %v", gotAccounts, tt.wantAccounts)
			}
		})
	}
}
