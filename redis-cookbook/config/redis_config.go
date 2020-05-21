package config

import (
	"time"
)

type RedisClusterConfig struct {
	NodeAddresses             []string
	PerNodePoolSize           int
	PerNodeMinIdleConnections int
	DialTimeout               time.Duration
	ReadTimeout               time.Duration
	WriteTimeout              time.Duration
}
