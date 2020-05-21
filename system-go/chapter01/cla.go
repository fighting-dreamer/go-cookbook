package chapter01

import (
	"fmt"
	"os"
)

func Start2() {
	args := os.Args
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}
