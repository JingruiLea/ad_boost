package main

import (
	"bytes"
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/lark"
	"github.com/JingruiLea/ad_boost/logic/auth"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
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

				// 从GET请求的查询字符串中获取参数
				queryParams := ctx.Request.URL.Query()
				for key, values := range queryParams {
					paramsMap[key] = values[0]
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
		Path:    "/api/v1/auth/oceanengine/callback",
		Handler: auth.OceanEngineCallback,
	})
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
