package ad

import (
	"context"
	"testing"
)

func TestGetSuggestBudget(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetSuggestBudget(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetSuggestBudget() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
