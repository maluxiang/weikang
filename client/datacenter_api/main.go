package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"weikang/client/datacenter_api/router"
	"weikang/client/datacenter_api/server"

	"weikang/start"
)

func init() {
	start.Init()
}

func main() {
	client, err := server.GetGRPCClient()
	if err != nil {
		zap.S().Error("服务器未启动", err)
		return
	}
	defer client.Close()
	r := gin.Default()
	r.Use(Cors())
	path := r.Group("/v1")
	router.DataCenter(path)
	r.Run(":8888")
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.GetHeader("Origin")

		if origin != "" {
			// 允许特定的域名访问
			c.Header("Access-Control-Allow-Origin", origin)
			// 允许的 HTTP 方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			// 允许的请求头
			c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization")
			// 允许携带凭证
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		// 处理 OPTIONS 方法的预检请求
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
