package chapter05

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFileWithBuffer(file *os.File) {
	buf := make([]byte, 8)
	_, err := io.ReadFull(file, buf)
	if err != nil {
		if err != io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}
	io.WriteString(os.Stdout, string(buf)) // only writing till the size of buffer bytes!!!
	fmt.Println()
}

func readfileUsingBifio(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Err() != nil {
			fmt.Println("error :", scanner.Err())
			os.Exit(1)
		}
		fmt.Println(line)
	}
}

func Start03() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("not enough arguments!!")
		os.Exit(1)
	}

	filename := args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	defer file.Close()

	//readFileWithBuffer(file)
	readfileUsingBifio(file)
}
