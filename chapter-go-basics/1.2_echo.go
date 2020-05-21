package main

import (
	"fmt"
	"os"
)

func main(){
	for it, arg := range os.Args {
		fmt.Println(it, arg)
	}
}
