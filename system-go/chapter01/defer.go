package chapter01

import "fmt"

// simple defer calls
func func1() {
	for i := 1; i <= 9; i++ {
		defer fmt.Print(i, " ")
	}
}

// defer on an anomymous no-argument func
// this makes func calls with using same local value at last time, theat is 10 here
//deferred value :  10
//deferred value :  10
//deferred value :  10
//deferred value :  10
//deferred value :  10
//deferred value :  10 ... so on
// NOTE: i gets incremented
func func2() {
	for i := 1; i <= 9; i++ {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}

// defer on an anomymous 1 argument func
// this makes func calls with i`s copies and gives result as the func-1
func func3() {
	for i := 1; i <= 9; i++ {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func Start5() {
	func1()
	fmt.Println()
	func2()
	fmt.Println()
	func3()
	fmt.Println()
}
