package wallet

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetWallet(ctx context.Context, adID int64) error {
	//   "list": [
	//      1748031128935424,
	//      1748031128935424,
	//      1767935594672136,
	//      1769126587284494
	//    ],

	params := map[string]interface{}{
		"advertiser_id": adID,
	}
	var resp GetWalletResp
	err := httpclient.NewClient().Get(ctx, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/finance/wallet/get/", &resp, params)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAdAccount httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetAdAccount respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}

type GetWalletRespData struct {
	TotalBalanceAbs int `json:"total_balance_abs"` //总余额

	BrandBalance              int `json:"brand_balance"`                 //品牌余额
	BrandBalanceValid         int `json:"brand_balance_valid"`           //品牌余额-可用余额
	BrandBalanceInvalid       int `json:"brand_balance_invalid"`         //品牌余额-不可用余额
	BrandBalanceInvalidFrozen int `json:"brand_balance_invalid_frozen"`  //品牌余额-不可用余额-冻结中
	BrandBalanceValidGrant    int `json:"brand_balance_valid_grant"`     //品牌余额-可用余额-赠款
	BrandBalanceValidNonGrant int `json:"brand_balance_valid_non_grant"` //品牌余额-可用余额-非赠款

	CommonValidGrantBalance int `json:"common_valid_grant_balance"` //赠款余额-巨量信息流广告-可用

	DeductionCouponBalance          int                   `json:"deduction_coupon_balance"`            //消返红包余额
	DeductionCouponBalanceAll       int                   `json:"deduction_coupon_balance_all"`        //消返红包余额（通用）
	DeductionCouponBalanceOther     int                   `json:"deduction_coupon_balance_other"`      //消返红包余额（代投）
	DeductionCouponBalanceSelf      int                   `json:"deduction_coupon_balance_self"`       //消返红包余额（自投）
	DefaultValidGrantBalance        int                   `json:"default_valid_grant_balance"`         //赠款余额-通用-可用
	GeneralBalanceInvalid           int                   `json:"general_balance_invalid"`             //通用余额-不可用余额
	GeneralBalanceInvalidFrozen     int                   `json:"general_balance_invalid_frozen"`      //通用余额-不可用余额-冻结
	GeneralBalanceInvalidOrder      int                   `json:"general_balance_invalid_order"`       //通用余额-不可用余额-随心推已下单
	GeneralBalanceValid             int                   `json:"general_balance_valid"`               //通用余额-可用余额
	GeneralBalanceValidGrantCommon  int                   `json:"general_balance_valid_grant_common"`  //通用余额-可用余额-赠款-巨量信息流广告
	GeneralBalanceValidGrantDefault int                   `json:"general_balance_valid_grant_default"` //通用余额-可用余额-赠款-通用
	GeneralBalanceValidGrantSearch  int                   `json:"general_balance_valid_grant_search"`  //通用余额-可用余额-赠款-巨量搜索广告
	GeneralBalanceValidGrantUnion   int                   `json:"general_balance_valid_grant_union"`   //通用余额-可用余额-赠款-穿山甲
	GeneralBalanceValidNonGrant     int                   `json:"general_balance_valid_non_grant"`     //通用余额-可用余额-非赠款
	GeneralTotalBalance             int                   `json:"general_total_balance"`               //通用余额
	GrantBalance                    int                   `json:"grant_balance"`                       //赠款余额
	GrantExpiring                   int                   `json:"grant_expiring"`                      //15天内赠款到期金额
	SearchValidGrantBalance         int                   `json:"search_valid_grant_balance"`          //赠款余额-巨量搜索广告-可用
	ShareExpiringDetailList         []ShareExpiringDetail `json:"share_expiring_detail_list"`          //共享赠款余额到期详情
	UnionValidGrantBalance          int                   `json:"union_valid_grant_balance"`           //赠款余额-穿山甲-可用
}

type ShareExpiringDetail struct {
	Category   string `json:"category"`
	Amount     int    `json:"amount"`
	ExpireTime int    `json:"expire_time"`
}

type GetWalletResp struct {
	ttypes.BaseResp
	Data GetWalletRespData `json:"data"`
}
