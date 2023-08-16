package boost

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/logic/ad"
	"github.com/JingruiLea/ad_boost/model/ttypes"
)

const (
	AdOnlineCount = 2
	AdRoiCount    = 5
	AdBidCount    = 3
)

type Roi float64

var roiList = []Roi{4, 4, 5, 5, 6, 6, 7, 7}

func Boost(ctx context.Context, adID int64, awemeID int64, adGroupID int64) error {
	var ad1 = ttypes.NewLiveCommonAd("test2", 2893532936291624, 1748031128935424, 1774203855472679).
		WithBudget(500).
		WithRoi(4)

	ad1.FirstIndustryId = 1904     //服装配饰
	ad1.SecondIndustryId = 190423  //男装
	ad1.ThirdIndustryId = 19042322 //休闲裤
	//获取建议Roi
	ecpRoi, err := ad.GetSuggestRoi(ctx, ad1)
	logs.CtxInfof(ctx, "计划建议Roi:%.2f", ecpRoi)

	////获取建议报价
	//lowBid, _, err := ad.GetSuggestBid(ctx, ad1)
	//if err != nil {
	//	logs.CtxErrorf(ctx, "GetSuggestBid error: %v", err)
	//	return err
	//}
	//ad1 = ad1.WithBid(lowBid) //先按最低报价出一下
	//创建广告
	adID, err = ad.CreateAd(ctx, ad1)
	if err != nil {
		logs.CtxErrorf(ctx, "CreateAd error: %v", err)
		return err
	}
	if err != nil {
		logs.CtxErrorf(ctx, "GetSuggestRoi error: %v", err)
		return err
	}
	return nil
}

func buildAudience(ctx context.Context, adID int64) error {
	//获取人群包
	//创建人群包
	//创建人群包定向
	return nil
}
