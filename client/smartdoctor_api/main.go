package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"weikang/client/smartdoctor_api/router"
	"weikang/client/smartdoctor_api/server"
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
	router.Doctor(path)
	r.Run(":8888")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
