package main

import (
	"fmt"
	"os"
	"time"
	"strings"
	"github.com/pkg/profile"
)

func withStringsJoin() int64 {
	startTime := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	endtime := time.Now()

	return endtime.Sub(startTime).Nanoseconds()
}

func simpleImplementation() int64 {
	startTime := time.Now()
	s, sep := "", ""
	for _, arg := range(os.Args) {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	endtime := time.Now()
	return endtime.Sub(startTime).Nanoseconds()
}

func main() {
	defer profile.Start().Stop()
	duration1 := simpleImplementation()
	duration2 := withStringsJoin()
	fmt.Println(duration1, duration2, duration1 > duration2)
}