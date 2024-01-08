package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Aweme struct {
	ID                      uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	AwemeAvatar             string         `json:"aweme_avatar"`
	AwemeHasLivePermission  bool           `json:"aweme_has_live_permission"`
	AwemeHasUniProm         bool           `json:"aweme_has_uni_prom"`
	AwemeHasVideoPermission bool           `json:"aweme_has_video_permission"`
	AwemeId                 int64          `json:"aweme_id"`
	AwemeName               string         `json:"aweme_name"`
	AwemeShowId             string         `json:"aweme_show_id"`
	AwemeStatus             string         `json:"aweme_status"`
	BindType                datatypes.JSON `json:"bind_type"`
	AccountID               int64          `json:"account_id"`

	gorm.Model
}

func (Aweme) TableName() string {
	return "aweme"
}
