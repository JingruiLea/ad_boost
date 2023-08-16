package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Audience struct {
	AudienceID int64          `json:"audience_id"`
	Name       string         `json:"name"`
	Config     datatypes.JSON `json:"config"`

	gorm.Model
}

func (a *Audience) TableName() string {
	return "audience"
}
