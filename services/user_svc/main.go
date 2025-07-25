package main

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"weikang/pkg"
	"weikang/services/user_svc/logic"
	"weikang/services/user_svc/proto/user"
	"weikang/services/user_svc/server"
	"weikang/start"
)

func init() {
	start.Init()
	err := pkg.InitRabbitMQ()
	if err != nil {
		panic("RabbitMQ连接失败: " + err.Error())
	}
}

func main() {
	listen, err := net.Listen("tcp", ":"+"8004")
	if err != nil {
		zap.S().Error("Listen error:", err)
		return
	}
	s := grpc.NewServer()
	user.RegisterUserServer(s, &logic.Server{})
	zap.S().Info("服务启动成功,端口:", listen.Addr())
	// 注册服务到Consul
	go func() {
		err = server.RegisterGRPCServerToConsul()
		if err != nil {
			zap.S().Error("服务注册到Consul失败:", err)
			return
		}
	}()
	if err = s.Serve(listen); err != nil {
		zap.S().Error("服务启动失败:", err)
	}
}
