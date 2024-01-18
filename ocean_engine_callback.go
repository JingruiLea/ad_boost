package main

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/logic/auth"
	"github.com/JingruiLea/ad_boost/logic/boost/sync"
	"github.com/JingruiLea/ad_boost/utils"
	"time"
)

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
	at, rt, err := auth.Auth(ctx, authCode)
	if err != nil {
		logs.CtxErrorf(ctx, "OceanEngineCallback Auth error. %s", err.Error())
		return nil, err
	}
	logs.CtxInfof(ctx, "got access_token: %s, refresh_token: %s", at, rt)
	err = sync.SyncAccount(ctx, at, rt)
	if err != nil {
		logs.CtxErrorf(ctx, "OceanEngineCallback sync.SyncAccount error. %s", err.Error())
		return nil, err
	}
	return "授权成功!", nil
}
