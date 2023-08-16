package ad

import (
	"context"
	"testing"
)

func TestUpdateAdRoiGoal(t *testing.T) {
	type args struct {
		ctx          context.Context
		advertiserID int64
		goals        []*RoiGoal
	}
	var tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.Background(),
				advertiserID: 1748031128935424,
				goals: []*RoiGoal{
					{
						AdId:    1774392648202275,
						RoiGoal: 7,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateAdRoiGoal(tt.args.ctx, tt.args.advertiserID, tt.args.goals); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdRoiGoal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
