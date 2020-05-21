package chapter05

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var mu sync.Mutex

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func WriteToFileWithLock(n int, file *os.File, w *sync.WaitGroup) {
	mu.Lock()
	time.Sleep(time.Duration(random(10, 1000)))
	fmt.Fprintf(file, "From %d writing %d\n", n, 2*n)
	fmt.Printf("Wrote to File %d\n", n)
	w.Done()
	mu.Unlock() // if you comment this out, you will get error from other go-routines!!!
}

func Start08() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Not enough arguments!!!")
		os.Exit(1)
	}

	fileName := args[1]
	number := 3
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}

	var w *sync.WaitGroup = new(sync.WaitGroup)
	w.Add(number)

	for i := 0; i < number; i++ {
		WriteToFileWithLock(i, file, w)
	}
	w.Wait()
}