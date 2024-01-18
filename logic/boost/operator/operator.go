package operator

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/lark"
	"github.com/JingruiLea/ad_boost/utils"
)

type AdOperator interface {
	StartAd(ctx context.Context, adID int64) error
	StopAd(ctx context.Context, adID int64) error
	AddBidTo(ctx context.Context, adID int64, target float64) error
	SubBidTo(ctx context.Context, adID int64, target float64) error
	AddBudgetTo(ctx context.Context, adID int64, target float64) error
	SubBudgetTo(ctx context.Context, adID int64, target float64) error
	CreateAd(ctx context.Context, createParams interface{}) error
	DeleteAd(ctx context.Context, adID int64) error
	CopyAd(ctx context.Context, adID int64) error
	AddRoiTargetTo(ctx context.Context, adID int64, target float64) error
	SubRoiTargetTo(ctx context.Context, adID int64, target float64) error
}

type FakeOperator struct {
}

func (f FakeOperator) AddBidTo(ctx context.Context, adID int64, target float64) error {
	lark.SendRoomMessage(ctx, fmt.Sprintf("修改出价: %d, %f", adID, target))
	return nil
}

func (f FakeOperator) SubBidTo(ctx context.Context, adID int64, target float64) error {
	lark.SendRoomMessage(ctx, fmt.Sprintf("修改出价: %d, %f", adID, target))
	return nil
}

func (f FakeOperator) AddBudgetTo(ctx context.Context, adID int64, target float64) error {
	lark.SendRoomMessage(ctx, fmt.Sprintf("修改预算: %d, %f", adID, target))
	return nil
}

func (f FakeOperator) SubBudgetTo(ctx context.Context, adID int64, target float64) error {
	lark.SendRoomMessage(ctx, fmt.Sprintf("修改预算: %d, %f", adID, target))
	return nil
}

func (f FakeOperator) CreateAd(ctx context.Context, createParams interface{}) error {
	lark.SendRoomMessage(ctx, fmt.Sprintf("创建计划: %s", utils.GetJsonStr(createParams)))
	return nil
}

func (f FakeOperator) DeleteAd(ctx context.Context, adID int64) error {
	lark.SendRoomMessage(ctx, fmt.Sprintf("删除计划: %d", adID))
	return nil
}

func (f FakeOperator) CopyAd(ctx context.Context, adID int64) error {
	lark.SendRoomMessage(ctx, fmt.Sprintf("复制计划: %d", adID))
	return nil
}

func (f FakeOperator) AddRoiTargetTo(ctx context.Context, adID int64, target float64) error {
	lark.SendRoomMessage(ctx, fmt.Sprintf("修改ROI目标: %d, %f", adID, target))
	return nil
}

func (f FakeOperator) SubRoiTargetTo(ctx context.Context, adID int64, target float64) error {
	lark.SendRoomMessage(ctx, fmt.Sprintf("修改ROI目标: %d, %f", adID, target))
	return nil
}

func (f FakeOperator) StartAd(ctx context.Context, adID int64) error {
	lark.SendRoomMessage(ctx, fmt.Sprintf("开始投放计划: %d", adID))
	return nil
}

func (f FakeOperator) StopAd(ctx context.Context, adID int64) error {
	lark.SendRoomMessage(ctx, fmt.Sprintf("暂停投放计划: %d", adID))
	return nil
}
