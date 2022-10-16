package config

import (
	"GooTask/global"
	"GooTask/logger"
	"GooTask/model/config"
	"github.com/spf13/viper"
)

func InitServerConfig() {
	// viper 实例化
	v := viper.New()
	// 设置文件路径
	v.SetConfigFile("./config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	// 设置初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	global.Settings = serverConfig
	logger.SugaredLogger.Info("Configuration loaded complete")
}
