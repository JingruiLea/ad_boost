package model

import (
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"gorm.io/gorm"
)

type Ad struct {
	AdID int64

	AdvertiserID    int64
	AdCreateTime    string
	AdModifyTime    string
	CampaignId      int64
	CampaignScene   string
	LabAdType       ttypes.LabAdType
	MarketingGoal   ttypes.MarketingGoal
	MarketingScene  ttypes.MarketingScene
	Name            string
	OptStatus       ttypes.OptStatus
	Status          ttypes.AdStatus
	DeliverySetting *DeliverySetting `gorm:"foreignKey:ID"`

	gorm.Model
}

func (*Ad) TableName() string {
	return "ad"
}

func (a *Ad) IsRoi() bool {
	return a.MarketingGoal == ttypes.MarketingGoalLivePromGoods &&
		a.DeliverySetting.ExternalAction == ttypes.ExternalActionAdConvertTypeLiveSuccessorderPay &&
		a.DeliverySetting.DeepExternalAction == ttypes.DeepExternalActionAdConvertTypeLivePayRoi &&
		a.DeliverySetting.DeepBidType == ttypes.DeepBidTypeMin
}

func (a *Ad) IsCpa() bool {
	return a.MarketingGoal == ttypes.MarketingGoalLivePromGoods &&
		a.DeliverySetting.ExternalAction == ttypes.ExternalActionAdConvertTypeLiveSuccessorderPay &&
		a.DeliverySetting.DeepExternalAction == ""
}
