package chapter03

import (
	"fmt"
	"runtime"
)

func Start03() {
	fmt.Print("we are using compiler : ", runtime.Compiler, " ")
	fmt.Print("on a machine : ", runtime.GOARCH, " ")
	fmt.Print("with a go version : ", runtime.Version(), " ")
	fmt.Print("using #goroutines : ", runtime.NumGoroutine())
}
