package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func withStringsJoin_org() int64 {
	startTime := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	endtime := time.Now()

	return endtime.Sub(startTime).Nanoseconds()
}

func simpleImplementation_org() int64 {
	startTime := time.Now()
	s, sep := "", ""
	for _, arg := range(os.Args[1:]) {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	endtime := time.Now()
	return endtime.Sub(startTime).Nanoseconds()
}

func main() {
	duration1 := simpleImplementation_org()
	duration2 := withStringsJoin_org()
	fmt.Println(duration1, duration2, duration1 > duration2)
}