package operator

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/logic/ad"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
)

type QianchuanOperator struct {
	AccountID int64
}

func (q QianchuanOperator) AddBidTo(ctx context.Context, adID int64, target float64) error {
	err := ad.UpdateAdBid(ctx, q.AccountID, []*ad.Bid{
		{
			AdId: adID,
			Bid:  utils.RoundFloat(target, 2),
		},
	})
	if err != nil {
		logs.CtxErrorf(ctx, "QianchuanOperator AddBidTo error: %v", err)
		return err
	}
	logs.CtxInfof(ctx, "QianchuanOperator AddBidTo success, adID: %d, target: %f", adID, target)
	return nil
}

func (q QianchuanOperator) SubBidTo(ctx context.Context, adID int64, target float64) error {
	err := ad.UpdateAdBid(ctx, q.AccountID, []*ad.Bid{
		{
			AdId: adID,
			Bid:  utils.RoundFloat(target, 2),
		},
	})
	if err != nil {
		logs.CtxErrorf(ctx, "QianchuanOperator SubBidTo error: %v", err)
		return err
	}
	logs.CtxInfof(ctx, "QianchuanOperator SubBidTo success, adID: %d, target: %f", adID, target)
	return nil
}

func (q QianchuanOperator) AddBudgetTo(ctx context.Context, adID int64, target float64) error {
	err := ad.UpdateAdBudget(ctx, q.AccountID, []*ad.Budget{
		{
			AdID:   adID,
			Budget: utils.RoundFloat(target, 2),
		},
	})
	if err != nil {
		logs.CtxErrorf(ctx, "QianchuanOperator AddBudgetTo error: %v", err)
		return err
	}
	logs.CtxInfof(ctx, "QianchuanOperator AddBudgetTo success, adID: %d, target: %f", adID, target)
	return nil
}

func (q QianchuanOperator) SubBudgetTo(ctx context.Context, adID int64, target float64) error {
	err := ad.UpdateAdBudget(ctx, q.AccountID, []*ad.Budget{
		{
			AdID:   adID,
			Budget: utils.RoundFloat(target, 2),
		},
	})
	if err != nil {
		logs.CtxErrorf(ctx, "QianchuanOperator SubBudgetTo error: %v", err)
		return err
	}
	logs.CtxInfof(ctx, "QianchuanOperator SubBudgetTo success, adID: %d, target: %f", adID, target)
	return nil
}

func (q QianchuanOperator) CreateAd(ctx context.Context, createParams interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (q QianchuanOperator) DeleteAd(ctx context.Context, adID int64) error {
	err := ad.UpdateAdStatus(ctx, &ad.UpdateAdStatusReq{
		AdvertiserID: q.AccountID,
		AdIDs:        []int64{adID},
		OptStatus:    ttypes.OptStatusDelete,
	})
	if err != nil {
		logs.CtxErrorf(ctx, "QianchuanOperator StopAd error: %v", err)
		return err
	}
	logs.CtxInfof(ctx, "QianchuanOperator StopAd success, adID: %d", adID)
	return nil
}

func (q QianchuanOperator) CopyAd(ctx context.Context, adID int64) error {
	//TODO implement me
	panic("implement me")
}

func (q QianchuanOperator) AddRoiTargetTo(ctx context.Context, adID int64, target float64) error {
	err := ad.UpdateAdRoiGoal(ctx, q.AccountID, []*ad.RoiGoal{
		{
			AdId:    adID,
			RoiGoal: target,
		},
	})
	if err != nil {
		logs.CtxErrorf(ctx, "ad.UpdateAdRoiGoal failed, err:%v", err)
		return err
	}
	logs.CtxInfof(ctx, "ad.UpdateAdRoiGoal success, adID: %d, target: %f", adID, target)
	return nil
}

func (q QianchuanOperator) SubRoiTargetTo(ctx context.Context, adID int64, target float64) error {
	err := ad.UpdateAdRoiGoal(ctx, q.AccountID, []*ad.RoiGoal{
		{
			AdId:    adID,
			RoiGoal: target,
		},
	})
	if err != nil {
		logs.CtxErrorf(ctx, "ad.UpdateAdRoiGoal failed, err:%v", err)
		return err
	}
	logs.CtxInfof(ctx, "ad.UpdateAdRoiGoal success, adID: %d, target: %f", adID, target)
	return nil
}

func (q QianchuanOperator) StartAd(ctx context.Context, adID int64) error {
	err := ad.UpdateAdStatus(ctx, &ad.UpdateAdStatusReq{
		AdvertiserID: q.AccountID,
		AdIDs:        []int64{adID},
		OptStatus:    ttypes.OptStatusEnable,
	})
	if err != nil {
		logs.CtxErrorf(ctx, "QianchuanOperator StartAd error: %v", err)
		return err
	}
	logs.CtxInfof(ctx, "QianchuanOperator StartAd success, adID: %d", adID)
	return nil
}

func (q QianchuanOperator) StopAd(ctx context.Context, adID int64) error {
	err := ad.UpdateAdStatus(ctx, &ad.UpdateAdStatusReq{
		AdvertiserID: q.AccountID,
		AdIDs:        []int64{adID},
		OptStatus:    ttypes.OptStatusDisable,
	})
	if err != nil {
		logs.CtxErrorf(ctx, "QianchuanOperator StopAd error: %v", err)
		return err
	}
	logs.CtxInfof(ctx, "QianchuanOperator StopAd success, adID: %d", adID)
	return nil
}
