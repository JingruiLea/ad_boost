package test_dal

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"testing"
)

func init() {
	dal.Init()
}

func TestCreateAweme(t *testing.T) {

	type args struct {
		ctx context.Context
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateAweme(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CreateAweme() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
