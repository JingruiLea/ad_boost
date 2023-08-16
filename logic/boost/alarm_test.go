package boost

import (
	"context"
	"testing"
	"time"
)

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
			report := GenReport(tt.args.ctx, tt.args.accountID)
			t.Log(report)
		})
	}
}
