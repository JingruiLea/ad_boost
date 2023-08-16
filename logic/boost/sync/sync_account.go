package sync

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/account_dal"
	"github.com/JingruiLea/ad_boost/dal/shop_dal"
	"github.com/JingruiLea/ad_boost/logic/account"
	"github.com/JingruiLea/ad_boost/model"
)

func SyncAccount(ctx context.Context, accessToken string, refreshToken string) error {
	//获取at对应的shops
	accounts, err := account.GetShopAccount(ctx, accessToken)
	if err != nil {
		logs.CtxErrorf(ctx, "SyncAccount account.GetShopAccount error: %v", err)
		return err
	}
	for _, a := range accounts {
		shop := buildShopFromAccount(a, accessToken, refreshToken)
		err = shop_dal.CreateOrUpdateShop(ctx, shop)
		if err != nil {
			logs.CtxErrorf(ctx, "SyncAccount shop_dal.CreateOrUpdateShop error: %v", err)
			return err
		}
		//从shop_id查出所有ad_id
		adIDs, err := account.GetAdAccountByShopID(ctx, shop.ShopID)
		if err != nil {
			logs.CtxErrorf(ctx, "SyncAccount account.GetAdAccountByShopID error: %v", err)
			return err
		}
		//根据ad_id查出所有ad_account
		adAccounts, err := account.MGetAdInfoDetail(ctx, adIDs)
		for _, adAccount := range adAccounts {
			err = account_dal.CreateOrUpdateAdAccount(ctx, adAccount.ToModel())
			if err != nil {
				logs.CtxErrorf(ctx, "SyncAccount shop_dal.CreateOrUpdateAdAccount error: %v", err)
				return err
			}
			//根据ad_account查出所有aweme_id
			awemeAccounts, err := account.GetAwemeByAdID(ctx, adAccount.ID, 1, 10)
			if err != nil {
				logs.CtxErrorf(ctx, "SyncAccount account.GetAwemeByAdID error: %v", err)
				return err
			}
			//TODO 大于10条的情况
			for _, info := range awemeAccounts.AwemeIdList {
				err = account_dal.CreateOrUpdateAweme(ctx, info.ToModel())
				if err != nil {
					logs.CtxErrorf(ctx, "SyncAccount shop_dal.CreateOrUpdateAweme error: %v", err)
					return err
				}
			}
		}
	}
	return nil
}

func buildShopFromAccount(acc *account.Account, at, rt string) *model.Shop {
	ret := &model.Shop{
		ID:           0,
		ShopID:       int64(acc.AdvertiserId),
		ShopName:     acc.AdvertiserName,
		AccountRole:  acc.AccountRole,
		AccessToken:  at,
		RefreshToken: rt,
	}
	if acc.IsValid {
		ret.IsValid = 1
	}
	return ret
}
