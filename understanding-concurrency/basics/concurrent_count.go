package basics

import (
	"fmt"
)

// infinetily printin index along with same input
func count(input string) {
	for i := 0; true; i++ {
		fmt.Println(i, input)
	}
}

// will only print sheep with count
func func1(){
	count("sheep")
}

// will only print sheep with count, the count("fist") never get executed
func func2() {
	count("sheep")
	go count("fish")
}

// will not print anything as by the time it start execution of count func, the func3() completes.
func func3() {
	go count("sheep")
	go count("fish")
}

func Start() {
	//func1()
	//func2()
	//func3()
}