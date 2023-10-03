package live_report

import (
	"context"
	"github.com/jinzhu/now"
	"testing"
	"time"
)

func TestGetReport(t *testing.T) {
	type args struct {
		ctx context.Context
		req *GetReportReq
	}
	fields := RoomMetricsFieldStatCost.All()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				req: NewGetReportReq(1748031128935424, 2893532936291624, now.New(time.Now().Add(time.Hour*48*-1)).BeginningOfDay(), now.New(time.Now().Add(time.Hour*48*-1)).EndOfDay(), fields),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GetReport(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("GetReport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
