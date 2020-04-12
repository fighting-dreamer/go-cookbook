package redis

import (
	"github.com/go-redis/redis/v7"
	"nipun.io/config"
)

func NewRedisClusterClient(redisClusterConfig config.RedisClusterConfig) *redis.ClusterClient {
	redisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        redisClusterConfig.NodeAddresses,
		DialTimeout:  redisClusterConfig.DialTimeout,
		ReadTimeout:  redisClusterConfig.ReadTimeout,
		WriteTimeout: redisClusterConfig.WriteTimeout,
		PoolSize:     redisClusterConfig.PerNodePoolSize,
		MinIdleConns: redisClusterConfig.PerNodeMinIdleConnections,
		//ReadOnly:     true,
	})
	return redisClient
}

func NewDefaultRedisClusterClient() *redis.ClusterClient {
	redisClusterConfig := config.GetRedisClusterConfig()
	redisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        redisClusterConfig.NodeAddresses,
		DialTimeout:  redisClusterConfig.DialTimeout,
		ReadTimeout:  redisClusterConfig.ReadTimeout,
		WriteTimeout: redisClusterConfig.WriteTimeout,
		PoolSize:     redisClusterConfig.PerNodePoolSize,
		MinIdleConns: redisClusterConfig.PerNodeMinIdleConnections,
		//ReadOnly:     true,
	})
	return redisClient
}

func NewDefaultRedisSimplerClient() *redis.Client {
	//redisConfig := config.GetRedisClusterConfig()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,  // use default DB
	})
	return redisClient
}
