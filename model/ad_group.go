package model

import (
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"gorm.io/gorm"
)

type AdGroup struct {
	AdGroupID      int64                 `json:"ad_group_id"`
	AccountID      int64                 `json:"account_id"`
	BudgetMode     ttypes.BudgetMode     `json:"budget_mode"`
	CreateDate     string                `json:"create_date"`
	MarketingGoal  ttypes.MarketingGoal  `json:"marketing_goal"`
	MarketingScene ttypes.MarketingScene `json:"marketing_scene"`
	Name           string                `json:"name"`
	Status         ttypes.AdGroupStatus  `json:"status"`
	Budget         float64               `json:"budget"`

	gorm.Model
}

func (*AdGroup) TableName() string {
	return "ad_group"
}

func (a *AdGroup) Display() utils.SortedList {
	return []*utils.KV{
		{Key: "id", Value: a.ID},
		{Key: "广告组名称", Value: a.Name},
		{Key: "广告组状态", Value: a.Status},
		{Key: "预算模式", Value: a.BudgetMode},
		{Key: "创建时间", Value: a.CreateDate},
		{Key: "营销目标", Value: a.MarketingGoal},
		{Key: "营销场景", Value: a.MarketingScene},
		{Key: "预算", Value: a.Budget},
	}
}
