package boost

import (
	"context"
	"github.com/JingruiLea/ad_boost/model/bo"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"testing"
)

func TestBoostOperator_Save(t *testing.T) {
	type fields struct {
		AdGroupID     int64
		AccountID     int64
		AwemeID       int64
		BoostParam    BoostParam
		RealTimeParam RealTimeParam
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{{
		name:   "",
		fields: fields{},
		args: args{
			ctx: context.Background(),
		},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ops[0]
			if err := op.Save(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBoostOperator_Load(t *testing.T) {
	type fields struct {
		AdGroupID     int64
		AccountID     int64
		AwemeID       int64
		BoostParam    BoostParam
		RealTimeParam RealTimeParam
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &BoostOperator{
				AccountID: 1784698853978186,
				AwemeID:   2691211639665967}
			if err := op.Load(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			print(utils.GetJsonStr(op))
		})
	}
}

func TestBoostOperator_GetAllAd(t *testing.T) {
	type fields struct {
		AdGroupID     int64
		AccountID     int64
		AwemeID       int64
		BoostParam    BoostParam
		RealTimeParam RealTimeParam
	}
	type args struct {
		ctx    context.Context
		status ttypes.AdStatus
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*bo.Ad
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx:    context.Background(),
				status: "",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ops[0]
			got, err := op.GetAllAd(tt.args.ctx, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllAd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			print(utils.GetJsonStr(got))
		})
	}
}
