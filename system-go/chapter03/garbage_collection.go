package chapter03

import (
	"fmt"
	"runtime"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem Alloc : ", mem.Alloc)
	fmt.Println("mem Total Alloc : ", mem.TotalAlloc)
	fmt.Println("mem Heap Alloc : ", mem.HeapAlloc)
	fmt.Println("mem Num GC : ", mem.NumGC)
}

func Start02() {
	for j := 0; j < 10; j++ {
		arr := make([]int, 1000000)
		for i := 0; i < 1000000; i++ {
			arr[i] = 100
		}
		printStats(runtime.MemStats{})
	}
	//look at "GODEBUG=gotrace=1"
}
