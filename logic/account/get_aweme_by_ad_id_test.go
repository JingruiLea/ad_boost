package account

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"testing"
)

func init() {
	redis_dal.Init()
}

func TestGetAwemeByAdID(t *testing.T) {
	type args struct {
		ctx  context.Context
		adID int64
		page int32
		size int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:  context.Background(),
				adID: 1748031128935424,
				page: 1,
				size: 10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GetAwemeByAdID(tt.args.ctx, tt.args.adID, tt.args.page, tt.args.size); (err != nil) != tt.wantErr {
				t.Errorf("GetAwemeByAdID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//aweme_id:2893532936291624
