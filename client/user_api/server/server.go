package server

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"weikang/client/user_api/proto/user"
	"weikang/global"
)

var (
	UserClient user.UserClient
)

func GetGRPCClient() (*grpc.ClientConn, error) {

	// 创建Consul客户端
	client, err := api.NewClient(&api.Config{Address: global.NacosConfig.Consul.Address})
	if err != nil {
		zap.S().Error("创建Consul客户端失败", err)
		return nil, err
	}

	// 从Consul中获取服务实例
	_, _, err = client.Health().Service(global.NacosConfig.Consul.RpcName, "", false, nil)
	if err != nil {
		zap.S().Error("从Consul中获取服务实例失败", err)
		return nil, err
	}

	// 选择一个服务实例

	grpcServerAddr := "127.0.0.1" + ":" + fmt.Sprintf("%d", 8004)
	conn, err := grpc.Dial(grpcServerAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		zap.S().Error("连接GRPC服务器失败", err)
		return nil, err
	}

	//创建一个grpc客户端并将其赋值给全局变量
	UserClient = user.NewUserClient(conn)

	return conn, nil
}
