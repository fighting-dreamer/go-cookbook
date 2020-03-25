package chapter04

import (
	"flag"
	"fmt"
)

func Start01() {
	minusO := flag.Bool("o", false, "set true OR false")
	minusI := flag.Int("i", 0, "set int value")
	minusK := flag.String("k", "empty-string", "set some string")

	flag.Parse()

	fmt.Println("-o :", *minusO)
	fmt.Println("-i :", *minusI)
	fmt.Println("-k :", *minusK)

	for index, val := range flag.Args() {
		fmt.Println(index, " : ", val)
	}

}
