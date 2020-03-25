package chapter02

import (
	"fmt"
	"strings"
)

func Start4() {
	var s [3]string
	s[0] = "1 2 3 4 5"
	s[1] = "-9 -8 -7 -6 -5 -4"
	s[2] = "10 12 14 13 11"
	column := 2
	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i])
		fmt.Println(data[column-1])
	}
}
