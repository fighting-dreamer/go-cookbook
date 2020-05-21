package chapter05

import (
	"fmt"
	"io/ioutil"
	"os"
)

func writeData(filename string, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0644)
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Read Error :", err)
		return "", err
	}
	aBytes := make([]byte, 100)
	n, err := file.Read(aBytes)
	fmt.Println("Read :", n, ", data :", string(aBytes))
	return string(aBytes) , err
}

func Start01() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please provide a FileName!!!")
		os.Exit(1)
	}
	filename := args[1]
	aByteSlice := "nipun Jindal"
	err := writeData(filename, aByteSlice)
	if err != nil {
		fmt.Println("Error :", err)
	}

	readFile(filename)
}
