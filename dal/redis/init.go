package redis

import (
	"fmt"
	"github.com/JingruiLea/ad_boost/config"
	"github.com/redis/go-redis/v9"
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
		panic("redis client init failed")
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}
