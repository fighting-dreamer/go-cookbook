package chapter01

import (
	"fmt"
	"os"
)

// stop when you get 0 as input
func stopWhenGetZero() {
	for {
		var x int
		fmt.Scanf("%d", &x)
		if x == 0 {
			break;
		}
		fmt.Println(x)
	}
}

func processCommandLineArgs() {
	args := os.Args

	minusK := false // minusK as false implies, we havent seen -k
	minusI := false // minusI as false implies, we havent seen -i

	for _, value := range args {
		if minusK == false && value == "-k " {
			minusK = true
		}
		if minusI == false && value == "-i" {
			minusI = true
		}
	}
	if minusK {
		fmt.Println("Got -k in the args!!!")
	}
	if minusI {
		fmt.Println("Got -i in the args!!!")
	}
}

func array2map (array []float64) map[int]float64 {
	mymap := make(map[int]float64)
	for index, value := range array {
		mymap[index] = value
	}
	return mymap
}

func createNewRandGenerator(seed int) func()int{
	return func()int {
		res := (seed * 3) % 1777777 + 1
		seed = res
		return res
	}
}

func array2slice(array [10]int)[]int {
	//res := make([]int, len(array))
	//for i, v := range array{
	//	res[i] = v
	//}
	return array[:]
}

func Start12() {
	//stopWhenGetZero()
	processCommandLineArgs()

	myMap := array2map([]float64{7.7, 7.9, 9.0, 9})
	fmt.Println(myMap)

	randGenerator := createNewRandGenerator(4)
	for i := 0; i < 20; i++ {
		fmt.Println(randGenerator())
	}

	x := array2slice([10]int{1,2,3,4,5,6,7,8,9,0})
	fmt.Println(x)

	// using copy :
	a := make([]int, 5)
	copy(a, x)
	fmt.Println(a)

	b := make([]int, 15)
	copy(b, a)
	fmt.Println(b)

	copy(b, x)
	fmt.Println(b)
}
