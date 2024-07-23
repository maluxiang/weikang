package start

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"weikang/global"
)

func Nacos() {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         global.LocalConfig.Nacos.SpaceId, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "../../tmp/nacos/log",
		CacheDir:            "../../tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      global.LocalConfig.Nacos.Address,
			ContextPath: "/nacos",
			Port:        uint64(global.LocalConfig.Nacos.Port),
			Scheme:      "http",
		},
	}
	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		zap.S().Error("创建动态配置客户端失败:", err.Error())
		return
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.LocalConfig.Nacos.DataId,
		Group:  global.LocalConfig.Nacos.Group})
	if err != nil {
		zap.S().Error("获取配置失败:", err.Error())
		return
	}
	err = json.Unmarshal([]byte(content), &global.NacosConfig)
	if err != nil {
		zap.S().Error("解析配置失败:", err.Error())
		return
	}
}
