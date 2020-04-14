package main

import (
	"nipun.io/config"
	"nipun.io/cookbook"
	"nipun.io/redis"
)

func main() {
	config.LoadConfig()
	redisClusterCli := redis.NewDefaultRedisClusterClient()
	redisCli := redis.NewDefaultRedisSimplerClient()

	cookbook.UseRedisAsRelationStore(redisClusterCli)
	cookbook.CreateUserNetwork(redisCli)
	cookbook.InvertedIndexesExample(redisCli)
	cookbook.InvertedIndexesWithScoreExample(redisCli)
	cookbook.RedisAsQueue(redisCli)

	redisCli.Close()
	redisClusterCli.Close()
}
