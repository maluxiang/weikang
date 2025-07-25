package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"weikang/client/medical_api/router"
	"weikang/client/medical_api/server"
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
	path := r.Group("/v1")
	router.Device(path)
	r.Run(":8885")
}
