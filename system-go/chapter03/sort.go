package chapter03

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type MyStruct struct {
	X int
	Y float64
	Z string
	A []int
}

func sort_non_decreasing(array []MyStruct) {
	sort.Slice(array, func(i, j int) bool {
		return array[i].X < array[j].X
	})
	fmt.Println(array)
}

func makeChange(array []int, len int) {
	index := rand.Intn(len)
	array[index] = rand.Intn(10000)
}

func concurrentOperationTestWhileSorting() {
	array := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		array[i] = rand.Intn(10000)
	}

	millisec := time.Millisecond

	c := make(chan int)
	go func() {
		sort.Slice(array, func(i, j int) bool {
			time.Sleep(millisec)
			return array[i] < array[j]
		})
		c <- 1
	}()
	time.Sleep(time.Nanosecond * 10)
	func() {
		for {
			select {
			case x := <-c:
				fmt.Println(x)
				return
			default:
				makeChange(array, 1000)

			}
		}
	}()

	fmt.Println(array) // this is not sorted, it implies sortin is not thread safe
}

func Start01() {
	array := []MyStruct{
		{X: 10,
			Y: 20.0,
			Z: "A",
			A: []int{1, 2, 3},
		},
		{X: 18,
			Y: 19.0,
			Z: "B",
			A: []int{4, 2, 3},
		},
		{X: 15,
			Y: 25.0,
			Z: "C",
			A: []int{1, 7, 3},
		},
	}
	fmt.Println(array)
	sort_non_decreasing(array)
	// lets see effect of concurency on sorting, is it thread safe!!!
	concurrentOperationTestWhileSorting()
}
