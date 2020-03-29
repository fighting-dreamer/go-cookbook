package chapter05

import (
	"fmt"
	"os"
	"strconv"
)

func Start07() {
	args := os.Args
	 if len(args) != 3 {
	 	fmt.Println("Not enough Arguments!!!")
	 	os.Exit(1)
	 }
	 SIZE, _ := strconv.ParseInt(args[1], 10, 64)
	 fileName := args[2]

	 _, err := os.Stat(fileName)
	 if err == nil {
	 	fmt.Println("file already exists")
	 	os.Exit(1)
	 }
	 fd, err := os.Create(fileName)
	 if err != nil {
	 	fmt.Println("error :", err)
	 }

	 _, err = fd.Seek(SIZE - 1, 0)
	 if err != nil {
	 	fmt.Println("error trying to seek :", err)
	 }
	 fd.Write([]byte{0})
	 if err != nil {
	 	fmt.Println("failed to write :", err)
	 }
	 err = fd.Close()
	 if err != nil {
	 	fmt.Println("error :", err)
	 }
}