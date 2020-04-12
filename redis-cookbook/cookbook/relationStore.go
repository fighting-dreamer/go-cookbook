package cookbook

import (
	"fmt"
	redisV7 "github.com/go-redis/redis/v7"
	"nipun.io/entity"
)

func userToMap(user entity.User) map[string]interface{} {
	addressString := fmt.Sprintf("%s",user.Address)
	userMap := make(map[string]interface{})
	userMap["NAME"] = user.Name
	userMap["ADDRESS"] = addressString
	userMap["AGE"] = user.Age
	userMap["WALLET"] = user.Wallet
	userMap["EMAIL"] = user.Email
	userMap["CONTACT"] = user.Contact
	userMap["TRANSAACTION_LIST_ID"] = user.TransactionListId

	return userMap
}

func createUserKey(userName string) string {
	userPrefix := "user:"
	userKey := userPrefix + userName
	return userKey
}

func setUserRelation(redisCli *redisV7.ClusterClient, user entity.User) {
	userKey := createUserKey(user.Name)
	userMap := userToMap(user)
	_, err := redisCli.HSet(userKey, userMap).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func UseRedisAsRelationStore(redisCli *redisV7.ClusterClient) {
	users := getUserData()

	// Setting users in Redis
	for _, user := range users {
		setUserRelation(redisCli, user)
	}

	// Getting users in Redis and printing them
	for _, user := range users {
		userKey := createUserKey(user.Name)
		result, _ := redisCli.HGet(userKey, "NAME").Result()
		fmt.Println("User Name :", result)

		resultAll, _ := redisCli.HGetAll(userKey).Result()
		fmt.Println("User :",resultAll)
		fmt.Println()
	}
}
