package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/url"
)

// QueryDecoder 创建一个中间件来解码所有的查询参数
func QueryDecoder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解码查询参数
		decodedParams := make(map[string][]string)
		for key, values := range c.Request.URL.Query() {
			for _, value := range values {
				decodedValue, err := url.QueryUnescape(value)
				if err != nil {
					// 错误处理
					c.AbortWithError(400, err)
					return
				}
				decodedParams[key] = append(decodedParams[key], decodedValue)
			}
		}

		// 创建一个新的URL对象，并设置解码后的查询参数
		newURL := *c.Request.URL
		newURL.RawQuery = url.Values(decodedParams).Encode()
		c.Request.URL = &newURL

		c.Next()
	}
}
