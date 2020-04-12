package config

type Config struct {
	redisClusterConfig RedisClusterConfig
}

var appConfig *Config

func GetRedisClusterConfig() RedisClusterConfig {
	return appConfig.redisClusterConfig
}
