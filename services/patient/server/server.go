package server

import (
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"weikang/global"
)

func RegisterGRPCServerToConsul() error {
	
	// 创建Consul客户端
	client, err := api.NewClient(&api.Config{Address: global.NacosConfig.Consul.Address})
	if err != nil {
		zap.S().Error("创建Consul客户端失败", err)
		return err
	}
	
	//定义Consul服务实例
	service := &api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    global.NacosConfig.Consul.RpcName,
		Tags:    nil,
		Port:    global.NacosConfig.Consul.RpcPortInt,
		Address: global.NacosConfig.Consul.RpcAddress,
	}
	
	// 注册服务实例
	err = client.Agent().ServiceRegister(service)
	if err != nil {
		zap.S().Error("注册服务实例失败", err)
		return err
	}
	
	return nil
}
