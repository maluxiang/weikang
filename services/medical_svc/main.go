package main

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"weikang/services/medical_svc/logic"
	"weikang/services/medical_svc/proto/medical"
	"weikang/services/medical_svc/server"
	"weikang/start"
)

func init() {
	start.Init()
}

func main() {
	listen, err := net.Listen("tcp", ":"+"8005")
	if err != nil {
		zap.S().Error("Listen error:", err)
		return
	}
	s := grpc.NewServer()
	medical.RegisterMedicalServiceServer(s, &logic.Server{})
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
