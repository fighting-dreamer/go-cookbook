package chapter01

import (
	"fmt"
	"math/rand"
)

// this will give error, as map is not concurrency safe.
func AddingElementsConcurrentlyInMap() {
	tempMap := make(map[int]int)
	for i := 0; i < 100000; i++ {
		go func() { tempMap[rand.Intn(10)] = rand.Intn(1000000) }()
	}
}

func tryAddingElementsConcurrentlyInMap() {
	tempMap := make(map[int]int, 10)
	for i := 0; i < 100000; i++ {
		go func() { tempMap[rand.Intn(10)] = rand.Intn(1000000) }()
	}
}

func Start9() {
	aMap := make(map[string]int)
	aMap["MON"] = 0
	aMap["TUE"] = 1
	aMap["WED"] = 2
	aMap["THU"] = 3
	aMap["FRI"] = 4
	aMap["SAT"] = 5
	aMap["SUN"] = 6

	fmt.Println("Sunday is : ", aMap["SUN"])

	_, ok := aMap["SUNDAY"]
	if ok {
		fmt.Println("Tuesday exists!!")
	} else {
		fmt.Println("Tuesday does not exists!!")
	}

	//iterating in map
	fmt.Println("iterating on the map")
	count := 0
	for key, _ := range aMap {
		count++
		fmt.Println(key)
	}
	fmt.Println("no. of elements in map : ", count)

	//deleting stuff in map
	delete(aMap, "FRI")    // remove the element from map
	delete(aMap, "FRIDAY") // does not return error if key is not found!!!

	//iterating on map
	fmt.Println("iterating on the map after delete")
	count = 0
	for key, _ := range aMap {
		count++
		fmt.Println(key)
	}
	fmt.Println("no. of elements in map : ", count)

	//add elements in map concurrently
	//AddingElementsConcurrentlyInMap()

	//tryin to fix the concurrency issue, when you know the size of map and assign space already
	tryAddingElementsConcurrentlyInMap() // does not work, need some kind of Locking.
}
