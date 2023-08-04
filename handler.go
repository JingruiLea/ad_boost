package main

import (
	"context"
	"github.com/JingruiLea/ad_boost/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinH struct {
	Method  string
	Path    string
	H       gin.HandlersChain
	Handler func(ctx context.Context, params map[string]string, requestBody []byte) (interface{}, error)
}

var handlers []*GinH

func Register(router *gin.Engine) {
	initRoutes()
	for _, handler := range handlers {
		handlerCopy := handler
		handlerCopy.H = []gin.HandlerFunc{
			func(ctx *gin.Context) {
				params := ctx.Params
				paramsMap := make(map[string]string)
				for _, param := range params {
					paramsMap[param.Key] = param.Value
				}
				requestBody, err := ctx.GetRawData()
				if err != nil {
					ctx.AbortWithStatus(500)
					return
				}
				result, err := handlerCopy.Handler(ctx, paramsMap, requestBody)
				if err != nil {
					ctx.AbortWithStatus(500)
					return
				}
				ctx.JSON(200, result)
			},
		}
		router.Handle(handlerCopy.Method, handlerCopy.Path, handlerCopy.H...)
	}
}

func Ping(ctx context.Context, params map[string]string, requestBody []byte) (interface{}, error) {
	return "pong", nil
}

func TestDB(ctx context.Context, params map[string]string, requestBody []byte) (interface{}, error) {
	return handler.TestDB(ctx, params, requestBody)
}

func TestRedis(ctx context.Context, params map[string]string, requestBody []byte) (interface{}, error) {
	return handler.TestRedis(ctx, params, requestBody)
}

func initRoutes() {
	handlers = append(handlers, &GinH{
		Method:  http.MethodGet,
		Path:    "/v1/ping",
		Handler: Ping,
	})
	handlers = append(handlers, &GinH{
		Method:  http.MethodGet,
		Path:    "/v1/test_db/:id",
		Handler: TestDB,
	})
	handlers = append(handlers, &GinH{
		Method:  http.MethodGet,
		Path:    "/v1/test_redis",
		Handler: TestRedis,
	})
}
