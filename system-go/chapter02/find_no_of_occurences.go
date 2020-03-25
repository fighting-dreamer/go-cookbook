package chapter02

import (
	"fmt"
	"strings"
)

func Start5() {
	var s [3]string
	s[0] = "1 2 3 4 5"
	s[1] = "-9 -8 -7 -6 -5 -4"
	s[2] = "1 1 4 3 12"
	countMap := make(map[string]int)

	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i])
		for _, value := range data {
			countMap[value]++
		}
	}
	fmt.Println(s) // this is : [1 2 3 4 5 -9 -8 -7 -6 -5 -4 1 1 4 3 12], Same as [15]int from appearence
	for k, v := range countMap {
		fmt.Println(k, " -> ", v)
	}
}
