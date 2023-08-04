package main

import (
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	Init()
	r := gin.Default()
	Register(r)
	err := r.Run("0.0.0.0:9000")
	if err != nil {
		panic(err)
	}
	logs.Infof("server started")
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	redis.Init()
}
