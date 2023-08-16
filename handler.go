package main

import (
	"bytes"
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/JingruiLea/ad_boost/lark"
	"github.com/JingruiLea/ad_boost/utils"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GinH struct {
	Method  string
	Path    string
	H       gin.HandlersChain
	Handler func(ctx context.Context, params map[string]string, requestBody []byte) (interface{}, error)
}

var handlers []*GinH

func Register(router *gin.RouterGroup) {
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
					logs.CtxErrorf(ctx, "handler error: %v", err)
					ctx.JSON(200, map[string]interface{}{
						"code": 500,
						"msg":  err.Error(),
						"data": nil,
					})
					return
				}
				ctx.JSON(200, map[string]interface{}{
					"code": 200,
					"msg":  "success",
					"data": result,
				})
			},
		}
		router.Handle(handlerCopy.Method, handlerCopy.Path, handlerCopy.H...)
	}
}

func Ping(ctx context.Context, params map[string]string, requestBody []byte) (interface{}, error) {
	return "pong", nil
}

func initRoutes() {
	handlers = append(handlers, &GinH{
		Method:  http.MethodGet,
		Path:    "/v1/ping",
		Handler: Ping,
	})
	//api/v1/oceanengine/callback
	handlers = append(handlers, &GinH{
		Method:  http.MethodGet,
		Path:    "/api/v1/oceanengine/callback",
		Handler: OceanEngineCallback,
	})
}

func OceanEngineCallback(ctx context.Context, params map[string]string, requestBody []byte) (interface{}, error) {
	logs.CtxInfof(ctx, utils.GetJsonStr(params))
	logs.CtxInfof(ctx, string(requestBody))
	authCode, ok := params["auth_code"]
	if !ok {
		return "no auto code", nil
	}
	err := redis_dal.GetRedisClient().Set(ctx, "auth_code", authCode, time.Hour*24).Err()
	if err != nil {
		logs.CtxErrorf(ctx, "OceanEngineCallback set redis_dal error. %s", err.Error())
		return nil, err
	}
	return "success", nil
}

func RegisterLark(router *gin.RouterGroup) {
	router.POST("/api/v1/lark/callback", func(ctx *gin.Context) {
		// Read and log request body
		requestBody, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			logs.CtxErrorf(ctx, "Error reading body: %v", err)
			return
		}
		logs.CtxDebugf(ctx, "Request body: %s", requestBody)

		// Since the body has been read, need to replace it for the handler
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		// Capture response body
		writer := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = writer

		// Call the handler
		lark.Register(ctx, ctx.Writer, ctx.Request)

		// Log response body
		logs.CtxDebugf(ctx, "Response body: %s", writer.body.String())
	})
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
