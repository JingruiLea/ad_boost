package model

import (
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"gorm.io/gorm"
)

type DeliverySetting struct {
	AdID               int64
	AdvertiserID       int64
	Budget             float64
	BudgetMode         ttypes.BudgetMode
	DeepBidType        ttypes.DeepBidType
	DeepExternalAction ttypes.DeepExternalAction
	EndTime            string
	ExternalAction     ttypes.ExternalAction
	ProductNewOpen     bool
	RoiGoal            float64
	SmartBidType       ttypes.SmartBidType
	StartTime          string
	CpaBid             float64

	gorm.Model
}

func (*DeliverySetting) TableName() string {
	return "delivery_setting"
}
