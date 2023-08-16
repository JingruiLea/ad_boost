package talent

import (
	"context"
	"testing"
)

func TestGetTalentByID(t *testing.T) {
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
				ctx:  context.Background(),
				adID: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetTalentByID(tt.args.ctx, tt.args.adID); (err != nil) != tt.wantErr {
				t.Errorf("GetTalentByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
