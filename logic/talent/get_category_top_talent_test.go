package talent

import (
	"context"
	"testing"
)

func TestGetCategoryTopTalent(t *testing.T) {
	type args struct {
		ctx        context.Context
		adID       int64
		categoryID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:        nil,
				adID:       0,
				categoryID: 13,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetCategoryTopTalent(tt.args.ctx, tt.args.adID, tt.args.categoryID); (err != nil) != tt.wantErr {
				t.Errorf("GetCategoryTopTalent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
