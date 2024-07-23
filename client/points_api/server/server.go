package server

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"weikang/client/points_api/proto/points"
	"weikang/global"
)

var (
	PointsClient points.PointsClient
)

func GetGRPCClient() (*grpc.ClientConn, error) {
	
	// 创建Consul客户端
	client, err := api.NewClient(&api.Config{Address: global.NacosConfig.Consul.Address})
	if err != nil {
		zap.S().Error("创建Consul客户端失败", err)
		return nil, err
	}
	
	// 从Consul中获取服务实例
	services, _, err := client.Health().Service(global.NacosConfig.Consul.RpcName, "", false, nil)
	if err != nil {
		zap.S().Error("从Consul中获取服务实例失败", err)
		return nil, err
	}
	
	// 选择一个服务实例
	service := services[0]
	grpcServerAddr := service.Service.Address + ":" + fmt.Sprintf("%d", service.Service.Port)
	conn, err := grpc.Dial(grpcServerAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		zap.S().Error("连接GRPC服务器失败", err)
		return nil, err
	}
	
	//创建一个grpc客户端并将其赋值给全局变量
	PointsClient = points.NewPointsClient(conn)
	
	return conn, nil
}
