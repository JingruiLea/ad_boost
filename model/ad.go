package model

import (
	"github.com/JingruiLea/ad_boost/model/bo"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"gorm.io/gorm"
)

type Ad struct {
	AdID           int64
	AdvertiserID   int64
	AdCreateTime   string
	AdModifyTime   string
	CampaignId     int64
	CampaignScene  ttypes.CampaignScene
	LabAdType      ttypes.LabAdType
	MarketingGoal  ttypes.MarketingGoal
	MarketingScene ttypes.MarketingScene
	Name           string
	OptStatus      ttypes.OptStatus
	Status         ttypes.AdStatus
	DeliverySetting
	AwemeInfo string
	Extra     string

	gorm.Model
}

func (*Ad) TableName() string {
	return "ad"
}

func (a *Ad) Display() utils.SortedList {
	ret := []*utils.KV{
		{Key: "id", Value: a.ID},
		{Key: "广告主ID", Value: a.AdvertiserID},
		{Key: "广告组ID", Value: a.CampaignId},
		{Key: "营销场景", Value: a.CampaignScene},
		{Key: "推广方式", Value: a.LabAdType},
		{Key: "营销目标", Value: a.MarketingGoal},
		{Key: "广告类型", Value: a.MarketingScene},
		{Key: "计划名称", Value: a.Name},
		{Key: "计划操作状态", Value: a.OptStatus},
		{Key: "计划投放状态", Value: a.Status},
		//{Key: "aweme_info", Value: a.AwemeInfo},
		{Key: "创建时间", Value: a.AdCreateTime},
		{Key: "修改时间", Value: a.AdModifyTime},
	}
	ret = append(ret, a.DeliverySetting.Display()...)
	return ret
}

func (*Ad) FromBO(a *bo.Ad) *Ad {
	ret := &Ad{
		AdCreateTime:   a.AdCreateTime,
		AdID:           a.AdID,
		AdModifyTime:   a.AdModifyTime,
		CampaignId:     a.CampaignId,
		CampaignScene:  a.CampaignScene,
		LabAdType:      a.LabAdType,
		MarketingGoal:  a.MarketingGoal,
		MarketingScene: a.MarketingScene,
		Name:           a.Name,
		OptStatus:      a.OptStatus,
		Status:         a.Status,
		DeliverySetting: DeliverySetting{
			SmartBidType:          a.DeliverySetting.SmartBidType,
			ExternalAction:        a.DeliverySetting.ExternalAction,
			DeepExternalAction:    a.DeliverySetting.DeepExternalAction,
			DeepBidType:           a.DeliverySetting.DeepBidType,
			ROIGoal:               a.DeliverySetting.ROIGoal,
			Budget:                a.DeliverySetting.Budget,
			ReviveBudget:          a.DeliverySetting.ReviveBudget,
			BudgetMode:            a.DeliverySetting.BudgetMode,
			CPABid:                a.DeliverySetting.CPABid,
			VideoScheduleType:     a.DeliverySetting.VideoScheduleType,
			LiveScheduleType:      a.DeliverySetting.LiveScheduleType,
			StartTime:             a.DeliverySetting.StartTime,
			EndTime:               a.DeliverySetting.EndTime,
			ScheduleTime:          a.DeliverySetting.ScheduleTime,
			ScheduleFixedRange:    a.DeliverySetting.ScheduleFixedRange,
			EnableAutoPause:       a.DeliverySetting.EnableAutoPause,
			AutoManageStrategyCmd: a.DeliverySetting.AutoManageStrategyCmd,
			EnableFollowMaterial:  a.DeliverySetting.EnableFollowMaterial,
			ProductNewOpen:        a.DeliverySetting.ProductNewOpen,
			QCPXMode:              a.DeliverySetting.QCPXMode,
			AllowQCPX:             a.DeliverySetting.AllowQCPX,
		},
	}
	ret.AwemeInfo = utils.GetJsonStr(a.AwemeInfo)
	return ret
}
