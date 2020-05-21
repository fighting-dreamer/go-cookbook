package chapter04

import (
	"fmt"
	"os"
)

func Start05() {
	args := os.Args
	file := args[0]

	fileInfo, err := os.Stat(file)
	if err == nil {
		fmt.Println(file, ":", fileInfo.Mode())
	}
}
