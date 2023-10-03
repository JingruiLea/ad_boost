package boost

import (
	"context"
	"testing"
)

func TestBoostInit(t *testing.T) {
	type args struct {
		ctx       context.Context
		accountID int64
	}
	var tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:       context.Background(),
				accountID: 1748031128935424,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BoostInit(tt.args.ctx, tt.args.accountID); (err != nil) != tt.wantErr {
				t.Errorf("BoostInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
