package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

type GetSuggestBudgetReq struct {
	AdvertiserID     int64                   `json:"advertiser_id"`
	AwemeID          int64                   `json:"aweme_id"`
	LiveScheduleType ttypes.LiveScheduleType `json:"live_schedule_type"`

	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`

	ScheduleTime string `json:"schedule_time,omitempty"`

	ScheduleFixedRange string `json:"schedule_fixed_range,omitempty"`
}

func GetSuggestBudget(ctx context.Context, req *GetSuggestBudgetReq) error {
	req.ScheduleTime = "111111111111111111111111111111111111111101111111111111111111111111111111111111111111111101111111111111111111111111111111111111111111111101111111111111111111111111111111111111111111111101111111111111111111111111111111111111111111111101111111111111111111111111111111111111111111111101111111111111111111111111111111111111111111111101111111"

	mmm := utils.Obj2Map(req)
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().AdGet(ctx, req.AdvertiserID, "https://api.oceanengine.com/open_api/v1.0/qianchuan/suggest/budget/", &resp, mmm)
	if err != nil {
		logs.CtxErrorf(ctx, "GetSuggestBudget httpclient.NewClient().Get error: %v", err)
		return err
	}
	fmt.Printf("GetSuggestBudget respMap: %s", utils.GetJsonStr(resp))
	//TODO Account
	return nil
}
