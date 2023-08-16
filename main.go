package main

import (
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/logic/boost"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	Init()
	r := gin.Default()

	adGroup := r.Group("/ad")

	Register(adGroup)
	RegisterLark(adGroup)

	err := r.Run("0.0.0.0:9000")
	if err != nil {
		panic(err)
	}
	log.Print("server started")
}

func Init() {
	boost.Init()
	redis_dal.Init()
	dal.Init()
}
