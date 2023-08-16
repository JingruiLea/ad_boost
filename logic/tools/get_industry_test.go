package tools

import (
	"context"
	"testing"
)

func TestGetIndustry1(t *testing.T) {
	type args struct {
		ctx   context.Context
		adID  int64
		level int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:   context.Background(),
				adID:  1748031128935424,
				level: 3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetIndustry(tt.args.ctx, tt.args.adID, tt.args.level); (err != nil) != tt.wantErr {
				t.Errorf("GetIndustry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
