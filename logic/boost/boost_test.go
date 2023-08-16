package boost

import (
	"context"
	"testing"
)

func TestBoost(t *testing.T) {
	type args struct {
		ctx     context.Context
		adID    int64
		awemeID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:     context.Background(),
				adID:    0,
				awemeID: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Boost(tt.args.ctx, tt.args.adID, tt.args.awemeID, 1774203855472679); (err != nil) != tt.wantErr {
				t.Errorf("Boost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
