package ad

import (
	"context"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"testing"
)

func TestGetSuggestRoi(t *testing.T) {
	type args struct {
		ctx context.Context
		ad  *ttypes.CreateAd
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
				ad:  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GetSuggestRoi(tt.args.ctx, tt.args.ad); (err != nil) != tt.wantErr {
				t.Errorf("GetSuggestRoi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
