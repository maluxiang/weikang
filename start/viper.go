package start

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"weikang/global"
)

func Viper() {
	viper.SetConfigFile("../../configs/configs.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		zap.S().Error("配置文件读取失败", err)
		return
	}
	err = viper.Unmarshal(&global.LocalConfig)
	if err != nil {
		zap.S().Error("配置文件解析失败", err)
		return
	}
}
