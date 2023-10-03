package redis_dal

import (
	"context"
	"fmt"
	"github.com/JingruiLea/ad_boost/config"
	"github.com/redis/go-redis/v9"
	"time"
)

var redisClient *redis.Client

func Init() {
	redisAddress := fmt.Sprintf("%s:%s", config.Configs.RedisHost, config.Configs.RedisPort)
	opt := &redis.Options{
		Addr:     redisAddress,
		Password: config.Configs.RedisPassword,
		DB:       config.GetRedisDBIndex(),
		Username: config.Configs.RedisUsername,
	}
	redisClient = redis.NewClient(opt)
	if redisClient == nil {
		panic("redis_dal client init failed")
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func Get(ctx context.Context, key string) (ret string, err error) {
	ret, err = redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		err = nil
	}
	return
}

func CheckDuplicateRequest(ctx context.Context, id string) (isDuplicate bool, err error) {
	// Check if the ID already exists in Redis.
	exists, err := redisClient.Exists(ctx, id).Result()
	if err != nil {
		return false, err
	}

	// If the ID exists, return true.
	if exists == 1 {
		return true, nil
	}

	// The ID does not exist, so add it to Redis.
	err = redisClient.Set(ctx, id, 1, 24*time.Hour).Err()
	if err != nil {
		return false, err
	}

	// The ID was added to Redis and did not previously exist.
	return false, nil
}
