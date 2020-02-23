package redis_related

import (
	"fmt"
	redis "github.com/go-redis/redis/v7"
	"strings"
	"time"
)

type RedisCli struct {
	Cli *redis.ClusterClient
}

func NewRedisCli() RedisCli {
	tenSeconds := time.Second*10

	redisClusterConfig := RedisClusterConfig{
		NodeAddresses:             strings.Split("localhost:30001,localhost:30002,localhost:30003,localhost:30004,localhost:30005,localhost:30006", ","),
		PerNodePoolSize:           100,
		PerNodeMinIdleConnections: 1,
		DialTimeout:               tenSeconds,
		ReadTimeout:               tenSeconds,
		WriteTimeout:              tenSeconds,
	}
	return RedisCli{
		Cli: redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        redisClusterConfig.NodeAddresses,
			DialTimeout:  redisClusterConfig.DialTimeout,
			ReadTimeout:  redisClusterConfig.ReadTimeout,
			WriteTimeout: redisClusterConfig.WriteTimeout,
			PoolSize:     redisClusterConfig.PerNodePoolSize,
			MinIdleConns: redisClusterConfig.PerNodeMinIdleConnections,
			ReadOnly:     true,
		})}
}

func (r RedisCli) Ping() string {
	res, err := r.Cli.Ping().Result()
	fmt.Println(res, err)
	return res
}

func (r RedisCli) Set(key string, value string) string {
	res, err := r.Cli.Set(key, value, 0).Result()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	//fmt.Println(res, err)
	return res
}


func (r RedisCli) Get(key string) string {
	res, err := r.Cli.Get(key).Result()
	if err != nil {
		return ""
	}
	return res
}

func (r RedisCli) MGet(keys []string) []interface{} {
	res, err := r.Cli.MGet(keys...).Result()
	if err != nil {
		// do nothing for now
	}
	return res
}

func (r RedisCli) Eval(script string, keys []string) interface{} {
	return r.Cli.Eval(script, keys, []string{"3"})
}

func (r RedisCli) Close() {
	r.Cli.Close()
}