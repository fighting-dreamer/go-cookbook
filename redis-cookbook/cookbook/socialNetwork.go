package cookbook

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func CreateUserNetwork(redisCli *redis.Client) {
	numbers := getNumberData()

	//Adding users as per there locality
	for _, number := range numbers {
		for _, setKey := range number.Categories {
			redisCli.SAdd(setKey, number.Value)
		}
	}
	categories := []string{odd, even, prime, perfectSquare}
	for _, setKey := range categories {
		numbers, _ := redisCli.SMembers(setKey).Result()
		fmt.Println("Set Key :", setKey, "Users :", numbers)
	}

	fmt.Println("BOTH PRIME and ODD : ", redisCli.SInter(prime, odd))
	fmt.Println("BOTH PRIME and EVEN : ", redisCli.SInter(prime, even))
	fmt.Println("BOTH ODD and EVEN : ", redisCli.SInter(odd, even))
	fmt.Println("BOTH ODD and PERFECT_SQUARE :", redisCli.SInter(odd, perfectSquare))
	fmt.Println("BOTH EVEN and PERFECT_SQUARE :", redisCli.SInter(even, perfectSquare))
}
