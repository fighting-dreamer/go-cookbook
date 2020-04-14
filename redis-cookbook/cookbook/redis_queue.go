package cookbook

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"math/rand"
	util "nipun.io/tools"
	"sync"
)

var wg sync.WaitGroup

func producer(redisCli *redis.Client, queueName string) {
	randInt := rand.Intn(1000)
	redisCli.RPush(queueName, randInt)
	wg.Done()
}

func consumer(redisCli *redis.Client, queueName string) string {
	result, err := redisCli.LPop(queueName).Result()
	if err != nil {
		return ""
	}
	fmt.Println("Consumed :", result)
	return result
}

func RedisAsQueue(redisCli *redis.Client) {
	queueName := util.GetUniqueID()
	fmt.Println(queueName)

	noOfProducer := 10
	noOfConsumer := 10
	wg.Add(noOfConsumer + noOfProducer)

	for i := 0; i < noOfProducer; i++ {
		go producer(redisCli, queueName)
	}
	for i := 0; i < noOfConsumer; i++ {
		go func() {
			result := consumer(redisCli, queueName)
			fmt.Println(result)
			wg.Done()
		}()
	}
	wg.Wait()
}
