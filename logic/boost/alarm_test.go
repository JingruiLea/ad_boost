package boost

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"testing"
	"time"
)

func init() {
	redis_dal.Init()
	dal.Init()
}

func TestAlarm(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Alarm(tt.args.ctx)
			time.Sleep(time.Minute * 100)
		})
	}
}

func TestGenReport(t *testing.T) {
	type args struct {
		ctx       context.Context
		accountID int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				ctx:       context.Background(),
				accountID: 1748031128935424,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			report := GenReportOnDelivery(tt.args.ctx, tt.args.accountID)
			t.Log(report)
		})
	}
}

func TestGenReportByRedisIDs(t *testing.T) {
	type args struct {
		ctx       context.Context
		accountID int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				ctx:       context.Background(),
				accountID: 1748031128935424,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenReportByRedisIDs(tt.args.ctx, tt.args.accountID); got != tt.want {
				t.Errorf("GenReportByRedisIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
