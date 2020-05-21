package chapter06

import (
	"fmt"
	"os"
)

func Start01() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Not enough Arguments!!!")
		os.Exit(1)
	}

	fileName := args[1]
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660) // open() is for read-only, this is a generic one!!!
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprintf(file, "Message : %s\n", "this is my message")
}
