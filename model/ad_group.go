package model

import "gorm.io/gorm"

type AdGroup struct {
	AdGroupID      int64   `json:"ad_group_id"`
	BudgetMode     string  `json:"budget_mode"`
	CreateDate     string  `json:"create_date"`
	MarketingGoal  string  `json:"marketing_goal"`
	MarketingScene string  `json:"marketing_scene"`
	Name           string  `json:"name"`
	Status         string  `json:"status"`
	Budget         float64 `json:"budget"`

	gorm.Model
}

func (*AdGroup) TableName() string {
	return "ad_group"
}
