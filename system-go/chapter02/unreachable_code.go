package chapter02

import "fmt"

func x() int {
	return -1
	// writing un-reachable code below!!
	fmt.Println("inside x()")
	return -10
}

func y() int {
	return -1
}
