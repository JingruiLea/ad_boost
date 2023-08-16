package ad_dal

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/model"
	"testing"
)

func init() {
	dal.Init()
}

func TestCreateOrUpdateDeliverySetting(t *testing.T) {
	type args struct {
		ctx             context.Context
		deliverySetting *model.DeliverySetting
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
				deliverySetting: &model.DeliverySetting{
					AdID:               123123,
					Budget:             110.12,
					BudgetMode:         "",
					DeepBidType:        "",
					DeepExternalAction: "",
					EndTime:            "",
					ExternalAction:     "",
					ProductNewOpen:     false,
					RoiGoal:            0,
					SmartBidType:       "",
					StartTime:          "",
					CpaBid:             0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateOrUpdateDeliverySetting(tt.args.ctx, tt.args.deliverySetting); (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdateDeliverySetting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
