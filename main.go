package main

import (
	"fmt"
	"github.com/JingruiLea/ad_boost/admin"
	"github.com/JingruiLea/ad_boost/admin/users"
	"github.com/JingruiLea/ad_boost/dal"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/lark"
	"github.com/JingruiLea/ad_boost/logic/boost"
	"github.com/JingruiLea/ad_boost/middlewares"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	Init()
	r := gin.Default()
	r.Use(middlewares.QueryDecoder())
	// JWT 中间件的设置
	authMiddleware, err := jwt.New(users.MyAuthenticator)

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/ad/admin/api/v1/login", authMiddleware.LoginHandler)
	r.POST("/ad/admin/api/v1/refresh_token", authMiddleware.RefreshHandler)

	adGroup := r.Group("/ad")
	adminGroup := adGroup.Group("/admin/api/v1/")

	Register(adGroup)
	RegisterLark(adGroup)
	admin.RegisterAdmin(adminGroup)

	err = r.Run("0.0.0.0:9000")
	if err != nil {
		panic(err)
	}
	log.Print("server started")
}

func Init() {
	lark.Init()
	//组件初始化
	redis_dal.Init()
	dal.Init()
	boost.InitOperators()
	//业务初始化
	boost.Init()
	fmt.Printf("init success. now:%s", time.Now().Format("2006-01-02 15:04:05"))
}
