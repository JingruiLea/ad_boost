package ad

import (
	"context"
	"testing"
)

func TestUpdateAdBudget(t *testing.T) {
	type args struct {
		ctx          context.Context
		advertiserID int64
		budgets      []*Budget
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.Background(),
				advertiserID: 1748031128935424,
				budgets: []*Budget{
					{
						AdId:   1774392648202275,
						Budget: 300.00,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateAdBudget(tt.args.ctx, tt.args.advertiserID, tt.args.budgets); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdBudget() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
