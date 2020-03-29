package chapter06

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//appending line number on file
func Start02() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Not enough arguments!!!")
		os.Exit(1)
	}

	fileName := args[1]
	lineNumber, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
	lines := strings.Split(string(fileContents), "\n")
	for i,line := range lines {
		lines[i] = fmt.Sprintf("%d : %s", lineNumber, line)
		lineNumber++
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0664)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}
}