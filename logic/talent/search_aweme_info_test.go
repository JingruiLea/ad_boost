package talent

import (
	"context"
	"testing"
)

func TestSearchAwemeInfo(t *testing.T) {
	type args struct {
		ctx     context.Context
		adID    int64
		keyWord string
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
				keyWord: "迪伦大叔品牌",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SearchAwemeInfo(tt.args.ctx, tt.args.adID, tt.args.keyWord); (err != nil) != tt.wantErr {
				t.Errorf("SearchAwemeInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
