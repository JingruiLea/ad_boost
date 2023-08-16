package model

import "gorm.io/gorm"

type Ad struct {
	AdID int64

	AdCreateTime    string
	AdModifyTime    string
	CampaignId      int64
	CampaignScene   string
	LabAdType       string
	MarketingGoal   string
	MarketingScene  string
	Name            string
	OptStatus       string
	Status          string
	DeliverySetting *DeliverySetting `gorm:"foreignKey:ID"`

	gorm.Model
}

func (*Ad) TableName() string {
	return "ad"
}
