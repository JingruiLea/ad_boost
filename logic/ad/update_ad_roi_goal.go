package ad

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/model/ttypes"
	"github.com/JingruiLea/ad_boost/utils"
	"github.com/JingruiLea/ad_boost/utils/httpclient"
)

func MUpdateAdRoiGoals(ctx context.Context, advertiserID int64, goals []*RoiGoal) error {
	var err error
	for _, goal := range goals {
		err = UpdateAdRoiGoal(ctx, advertiserID, []*RoiGoal{goal})
		if err != nil {
			logs.CtxErrorf(ctx, "MUpdateAdRoiGoals UpdateAdRoiGoal error: %v", err)
			return err
		}
	}
	return nil
}

func UpdateAdRoiGoal(ctx context.Context, advertiserID int64, goals []*RoiGoal) error {
	var req UpdateAdRoiGoalReq
	req.AdvertiserId = advertiserID
	req.RoiGoalUpdates = goals
	var resp = make(map[string]interface{})
	err := httpclient.NewClient().AdPost(ctx, advertiserID, "https://ad.oceanengine.com/open_api/v1.0/qianchuan/roi/goal/update", req, &resp)
	if err != nil {
		logs.CtxErrorf(ctx, "UpdateAdRoiGoal httpclient.NewClient().Post error: %v", err)
		return err
	}
	fmt.Printf("UpdateAdRoiGoal respMap: %s", utils.GetJsonStr(resp))
	return nil
}

type UpdateAdRoiGoalReq struct {
	AdvertiserId   int64      `json:"advertiser_id"`
	RoiGoalUpdates []*RoiGoal `json:"roi_goal_updates"`
}

type RoiGoal struct {
	AdId    int64   `json:"ad_id"`
	RoiGoal float64 `json:"roi_goal"`
}

type UpdateAdRoiGoalResp struct {
	Data *UpdateAdRoiGoalRespData `json:"data"`
	ttypes.BaseResp
}

type UpdateAdRoiGoalRespData struct {
	Results []*UpdateAdRoiGoalRespDataResults `json:"results"`
}

type UpdateAdRoiGoalRespDataResults struct {
	AdId         int64  `json:"ad_id"`
	ErrorMessage string `json:"error_message"`
	Flag         bool   `json:"flag"`
}
