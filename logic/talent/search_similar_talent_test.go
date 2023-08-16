package talent

import (
	"context"
	"testing"
)

func TestSearchSimilarTalent(t *testing.T) {
	type args struct {
		ctx         context.Context
		adID        int64
		awemeShowID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:         nil,
				adID:        1748031128935424,
				awemeShowID: "3227791139284568",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SearchSimilarTalent(tt.args.ctx, tt.args.adID, tt.args.awemeShowID); (err != nil) != tt.wantErr {
				t.Errorf("SearchSimilarTalent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
