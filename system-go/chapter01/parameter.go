package chapter01

import (
	"fmt"
	"os"
	"strings"
)

func Start4() {
	args := os.Args
	minusI := false

	for i := 0; i < len(args); i++ {
		if (strings.Compare(args[i], "-i") == 0) {
			minusI = true;
			break;
		}
	}
	if(minusI) {
		fmt.Println("Got the -i parameter!!!")
		fmt.Print("y/n: ")
		var answer string
		fmt.Scanln(&answer)
		fmt.Println("Got answer: ", answer)
	}else {
		fmt.Println("-i paramenter is not set")
	}
}

