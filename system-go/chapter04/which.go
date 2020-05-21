package chapter04

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func testIfFileExists(filepath string) bool {
	fileInfo, err := os.Stat(filepath)
	fmt.Println(filepath)
	if err == nil {
		mode := fileInfo.Mode()
		if mode.IsRegular() {
			if mode&0111 != 0 {
				return true
			}
		}
	}
	return false
}

func Start04() {
	minusA := flag.Bool("a", false, "a")
	minusS := flag.Bool("s", false, "s")
	flag.Parse()

	flags := flag.Args()
	if len(flags) == 0 {
		fmt.Println("Please provide input")
		os.Exit(1)
	}
	filename := flags[0]
	foundIt := false

	path := os.Getenv("PATH")
	pathSlice := strings.Split(path, ":")

	for _, directory := range pathSlice {
		fullPath := directory + "/" + filename
		foundIt = testIfFileExists(fullPath)
		if foundIt {
			if *minusS == true {
				os.Exit(0)
			}
			fmt.Println(fullPath)
			if *minusA != true {
				os.Exit(0)
			}
		}
	}
	if foundIt == false {
		os.Exit(1)
	}
}
