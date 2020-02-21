package basics

import (
	"fmt"
	"time"
	"sync"
)

// infinetily printin index along with same input
func count(input string) {
	for i := 0; true; i++ {
		fmt.Println(i, input)
	}
}

// infinetily printin index along with same input
func count2(input string) {
	halfSecond := time.Millisecond*500
	for i := 0; i <= 5; i++ {
		fmt.Println(i, input)
		time.Sleep(halfSecond)
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


// this is blocking, will go-routine completes, here it will run infinitely
func func4() {
	var wg sync.WaitGroup
	wg.Add(1) // imlies , I am waiting for '1' go-routine to complete

	go func() {
		count("sheep")
		wg.Done()
	}() // anonymous function, coz, the wait-group related changes are not responsibility of go-routine
	wg.Wait()
}

// this is blocking, will go-routine completes, here it will run infinitely
func func5() {
	var wg sync.WaitGroup
	wg.Add(1) // imlies , I am waiting for '1' go-routine to complete

	go func() {
		count2("sheep")
		wg.Done()
	}() // anonymous function, coz, the wait-group related changes are not responsibility of go-routine
	wg.Wait()
}


func Start() {
	//func1()
	//func2()
	//func3()

	//With Waitgroup
	//func4()
	func5()
}