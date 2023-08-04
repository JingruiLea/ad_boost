package handler

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/redis"
	"github.com/JingruiLea/ad_boost/dal/test_dal"
	"github.com/JingruiLea/ad_boost/utils"
)

func TestDB(ctx context.Context, params map[string]string, body []byte) (interface{}, error) {
	idStr, ok := params["id"]
	if !ok {
		return nil, nil
	}
	id := utils.Str2I64(idStr, 0)
	data, err := test_dal.GetTestByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func TestRedis(ctx context.Context, params map[string]string, body []byte) (interface{}, error) {
	client := redis.GetRedisClient()
	err := client.Set(ctx, "test", "test", 0).Err()
	if err != nil {
		logs.CtxErrorf(ctx, "redis set failed: %v", err)
		return nil, err
	}
	val, err := client.Get(ctx, "test").Result()
	if err != nil {
		logs.CtxErrorf(ctx, "redis get failed: %v", err)
		return nil, err
	}
	return val, nil
}
