package account

import (
	"context"
	"testing"
)

func TestGetAdInfoDetail(t *testing.T) {
	type args struct {
		ctx   context.Context
		adIDs []int64
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
				adIDs: []int64{1748031128935424},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := MGetAdInfoDetail(tt.args.ctx, tt.args.adIDs); (err != nil) != tt.wantErr {
				t.Errorf("MGetAdInfoDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
