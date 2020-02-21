package main

import (
	"fmt"
	"os"
	"time"
)

func my_prints(start time.Time) {
	fmt.Println(os.Getpid())
	for j := 0; j <= 100000000; j++ {
		go time.Sleep(time.Millisecond)
		fmt.Println(time.Now().Sub(start).Seconds())
	}
}

func main() {
	go my_prints(time.Now())
	time.Sleep(time.Second * 5)
}
