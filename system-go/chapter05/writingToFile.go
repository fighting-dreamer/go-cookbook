package chapter05

import (
	"fmt"
	"os"
)

func Start04() {
	args := os.Args

	if (len(args) != 2) {
		fmt.Println("Not Enough input!!!")
		os.Exit(1)
	}

	filename := args[1]
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("got error :", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprintf(file, "[%s]", "temp")
	fmt.Fprintf(file, "Using Fprintf to write stuff\n") // it rites extra string = temp
}
