package wallet

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"testing"
)

func init() {
	dal.Init()
	redis_dal.Init()
}

func TestGetWallet(t *testing.T) {
	type args struct {
		ctx  context.Context
		adID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				//1785607285728259
				//1777522841147527
				//1769643406656584
				//1786246903614537
				//1784698853978186 6152 05000
				//1786151003464714 43525 88000
				//1786309327331339 8020 76000
				//1703886601680909 30984 30000
				//1765679469538382
				//1772449904394254 7326 81000
				ctx:  context.Background(),
				adID: 1772449904394254,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GetWallet(tt.args.ctx, tt.args.adID); (err != nil) != tt.wantErr {
				t.Errorf("GetWallet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
