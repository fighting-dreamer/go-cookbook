package chapter04

import (
	"fmt"
	"os"
)

func Start06() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide the arguments")
		os.Exit(1)
	}
	file := args[1]
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
		return
	}
}
