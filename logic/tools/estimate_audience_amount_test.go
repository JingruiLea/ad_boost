package tools

import (
	"context"
	"reflect"
	"testing"
)

func TestEstimateAudienceAmountByAudienceID(t *testing.T) {
	type args struct {
		ctx        context.Context
		audienceID int64
		awemeID    int64
	}
	tests := []struct {
		name    string
		args    args
		want    *EstimateAudienceAmountData
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:        context.Background(),
				audienceID: 7320545678319681555,
				awemeID:    2691211639665967,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EstimateAudienceAmountByAudienceID(tt.args.ctx, tt.args.audienceID, tt.args.awemeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("EstimateAudienceAmountByAudienceID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EstimateAudienceAmountByAudienceID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
