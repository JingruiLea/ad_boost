package model

import (
	"github.com/JingruiLea/ad_boost/model/bo"
	"github.com/JingruiLea/ad_boost/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Audience struct {
	AudienceID int64          `json:"audience_id"`
	Name       string         `json:"name"`
	Config     datatypes.JSON `json:"config"`
	AccountID  int64          `json:"account_id"`

	gorm.Model
}

func (a *Audience) TableName() string {
	return "audience"
}

func (*Audience) FromBO(a *bo.Audience, accountID int64) *Audience {
	ret := &Audience{
		AudienceID: a.OrientationID,
		Name:       a.OrientationName,
		Config:     datatypes.JSON(utils.GetJsonStr(a)),
		AccountID:  accountID,
	}
	return ret
}
