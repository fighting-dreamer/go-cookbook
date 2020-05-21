package chapter05

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
)

func Start02() {
	args := os.Args
	input := args[1]
	aNumber, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, aNumber)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	fmt.Printf("Little Endian : number : %d Buffer : %x\n", aNumber, buf)

	buf.Reset()

	err = binary.Write(buf, binary.BigEndian, aNumber)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	fmt.Printf("Big Endian : number : %d Buffer : %x", aNumber, buf)
}