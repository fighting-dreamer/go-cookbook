package config

import (
	"github.com/spf13/viper"
	util "nipun.io/tools"
	"strings"
	"time"
)

func loadConfigFromLocalFile() {
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()
}

func LoadConfig() {
	loadConfigFromLocalFile()
	timeOut := time.Duration(util.GetIntOrPanic(RedisDefaultTimeOutSeconds)) * time.Second
	appConfig = &Config{
		redisClusterConfig: RedisClusterConfig{
			NodeAddresses:             strings.Split(util.FatalGetString(RedisClusterNodeAddresses), ","),
			PerNodePoolSize:           util.GetIntOrPanic(RedisClusterPerNodePooLSize),
			PerNodeMinIdleConnections: util.GetIntOrPanic(RedisClusterPerNodeIdleConnSize),
			DialTimeout:               timeOut,
			ReadTimeout:               timeOut,
			WriteTimeout:              timeOut,
		},
	}
}
