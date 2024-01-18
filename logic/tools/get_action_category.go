package tools

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func GetActionCategory(ctx context.Context, req *GetActionCategoryReq) (*GetActionCategoryData, error) {
	var resp GetActionCategoryData
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://ad.oceanengine.com/open_api/2/tools/interest_action/action/category/", &resp, utils.Obj2Map(req))
	if err != nil {
		logs.CtxErrorf(ctx, "GetCategoryList httpclient.NewClient().Get error: %v", err)
		return nil, err
	}
	fmt.Printf("GetCategoryList respMap: %s", utils.GetJsonStr(resp))
	return &resp, err
}

type GetActionCategoryReq struct {
	AdvertiserID int64                `json:"advertiser_id"` // 广告主ID
	ActionScene  []ttypes.ActionScene `json:"action_scene"`  // 行为场景
	ActionDays   ttypes.ActionDays    `json:"action_days"`   // 行为天数
}

// BehaviorCategory 表示行为类目的结构体
type BehaviorCategory struct {
	ID       string              `json:"id"`                 // 类目ID
	Name     string              `json:"name"`               // 类目名称
	Children []*BehaviorCategory `json:"children,omitempty"` // 子类目列表
	Num      string              `json:"num"`                // 数量
}

// Data 表示JSON返回值的顶层结构
type GetActionCategoryData []*BehaviorCategory
