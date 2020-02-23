package redis_related

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

const prefix = "temp-"

func setDataInRedis(rc *RedisCli) {
	for i := 0; i < 100000000; i++ {
		//fmt.Println(prefix + strconv.Itoa(i), prefix + strconv.Itoa(rand.Int()))
		rc.Set(prefix + strconv.Itoa(i), prefix + strconv.Itoa(rand.Int()))
	}
}

func setDataInRedisWithHashTags(rc *RedisCli) {
	for i := 0; i < 5000000; i++ {
		//fmt.Println(prefix + strconv.Itoa(i), prefix + strconv.Itoa(rand.Int()))
		rc.Set(prefix + "{" + strconv.Itoa(rand.Intn(24)) + "}" +  strconv.Itoa(i), prefix + strconv.Itoa(rand.Int()))
	}
}

// Get on multiple keys
func getOnMultipleKeysSerial(rc *RedisCli) {
	keys := []string{}
	for i := 0; i < 10; i++ {
		keys = append(keys, prefix + strconv.Itoa(rand.Intn(5000000)))
	}

	for _, value := range keys {
		rc.Get(value)
	}
}

//Get on multiple keys, goroutine
func getOnMultipleKeysConncurent(rc *RedisCli) {
	keys := []string{}
	for i := 0; i < 10; i++ {
		keys = append(keys, prefix + strconv.Itoa(rand.Intn(5000000)))
	}
	var wg sync.WaitGroup
	wg.Add(len(keys))
	for _, value := range keys {
		go func() {
			defer wg.Done()
			rc.Get(value)
		}()
	}
	wg.Wait()
}

func mgetOnKeys(rc *RedisCli) {
	keys := []string{}
	for i := 0; i < 10; i++ {
		keys = append(keys, prefix + strconv.Itoa(rand.Intn(5000000)))
	}
	rc.MGet(keys)
}

func evalOnKeys(rc *RedisCli) {
	keys := []string{}
	for i := 0; i < 10; i++ {
		keys = append(keys, prefix + "{3}" + strconv.Itoa(rand.Intn(5000000)))
	}
	script := "return {redis.call('get', 'qwerty1'), redis.call('get', 'qwerty2'), redis.call('get', 'qwerty3'), redis.call('get', 'qwerty4'), redis.call('get', 'qwerty5'), redis.call('get', 'qwerty0', 1)} 0"
	getCallString := "redis.call('get', '%s')"
	finalGetCallString := fmt.Sprintf(getCallString, keys[0])
	for _, value := range(keys[1:]) {
		finalGetCallString += "," + fmt.Sprintf(getCallString, value)
	}
	//fmt.Println(finalGetCallString)
	script = "return {" + finalGetCallString + "}"
	//fmt.Println()
	fmt.Println(rc.Eval(script, []string{}))
}

func Start() {
	rediscli := NewRedisCli()
	fmt.Println(rediscli.Ping())
	//setDataInRedis(&rediscli)
	setDataInRedisWithHashTags(&rediscli)
	//evalOnKeys(&rediscli)
	rediscli.Close()
}
