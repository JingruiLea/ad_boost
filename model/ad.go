package model

import (
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"gorm.io/gorm"
)

type Ad struct {
	AdID           int64
	AdvertiserID   int64
	AdCreateTime   string
	AdModifyTime   string
	CampaignId     int64
	CampaignScene  string
	LabAdType      ttypes.LabAdType
	MarketingGoal  ttypes.MarketingGoal
	MarketingScene ttypes.MarketingScene
	Name           string
	OptStatus      ttypes.OptStatus
	Status         ttypes.AdStatus
	DeliverySetting

	gorm.Model
}

func (*Ad) TableName() string {
	return "ad"
}
