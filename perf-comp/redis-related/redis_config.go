package redis_related

import "time"

type RedisClusterConfig struct {
	NodeAddresses             []string
	PerNodePoolSize           int
	PerNodeMinIdleConnections int
	DialTimeout               time.Duration
	ReadTimeout               time.Duration
	WriteTimeout              time.Duration
}

func NewRedisClusterConfig(nodeAddresses []string, perNodePoolSize int, perNodeMinIdleConnections int, dialTimeout time.Duration, readTimeout time.Duration, writeTimeout time.Duration) RedisClusterConfig {
	return RedisClusterConfig{
		NodeAddresses:             nodeAddresses,
		PerNodePoolSize:           perNodePoolSize,
		PerNodeMinIdleConnections: perNodeMinIdleConnections,
		DialTimeout:               dialTimeout,
		ReadTimeout:               readTimeout,
		WriteTimeout:              writeTimeout,
	}
}
