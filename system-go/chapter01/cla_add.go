package chapter01

import (
	"fmt"
	"os"
	"strconv"
)

func Start3() {
	args := os.Args
	sum := 0
	for i := 1; i < len(args); i++ {
		temp, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Println(err.Error())
		}
		sum += temp
	}
	fmt.Println("Sum: ", sum)
}
