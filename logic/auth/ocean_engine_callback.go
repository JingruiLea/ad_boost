package auth

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/utils"
	"time"
)

const InitAccountID = 1008611

func OceanEngineCallback(ctx context.Context, params map[string]string, requestBody []byte) (interface{}, error) {
	logs.CtxInfof(ctx, utils.GetJsonStr(params))
	logs.CtxInfof(ctx, string(requestBody))
	authCode, ok := params["auth_code"]
	if !ok {
		return "no auth code", nil
	}
	err := redis_dal.GetRedisClient().Set(ctx, "auth_code", authCode, time.Hour*24).Err()
	if err != nil {
		logs.CtxErrorf(ctx, "OceanEngineCallback set redis_dal error. %s", err.Error())
		return nil, err
	}
	at, rt, err := Auth(ctx, authCode)
	if err != nil {
		logs.CtxErrorf(ctx, "OceanEngineCallback Auth error. %s", err.Error())
		return nil, err
	}
	logs.CtxInfof(ctx, "got access_token: %s, refresh_token: %s", at, rt)
	err = SaveAtRtToRedis(ctx, at, rt, InitAccountID)
	if err != nil {
		logs.CtxErrorf(ctx, "OceanEngineCallback SaveAtRtToRedis error. %s", err.Error())
		return nil, err
	}
	return "授权成功!", nil
}
