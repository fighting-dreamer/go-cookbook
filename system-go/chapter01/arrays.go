package chapter01

import (
	"fmt"
	"math/rand"
	"time"
)

func havingFunWithIndexIssues() {
	myArray := [4]int{1, 2, 3, 4}
	my2dArray := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	my3dArray := [][][]int{{{1, 2}, {3, 4}, {4, 5}}, {{6, 7}, {8, 9}, {10, 11}}, {{12, 13}, {14, 15}, {16, 17}}}

	// getting Error when trying to get values at wrong indexes
	//fmt.Println(myArray[-1]) // go wont compile this!!! Do it like below
	i := -1
	fmt.Println(myArray[i]) // go wont compile this!!!
	fmt.Println(my2dArray[1][10])
	fmt.Println(my3dArray[2][10][1])
}

func reWritingStackWithNegtiveIndices() {
	myArray := [4]int{1, 2, 3, 4}
	my2dArray := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	my3dArray := [][][]int{{{1, 2}, {3, 4}, {4, 5}}, {{6, 7}, {8, 9}, {10, 11}}, {{12, 13}, {14, 15}, {16, 17}}}

	i := -2
	myArray[i] = -9
	fmt.Println(myArray[i])
	i = -1
	j := -2
	my2dArray[i][j] = -3
	fmt.Println(my2dArray[i][j])
	i = -2
	j = -1
	k := -2
	my3dArray[i][j][k] = -4
	fmt.Println(my3dArray[i][j][k])
}

func printMyArray(array []int) {
	for _, value := range array {
		fmt.Println(value)
	}
}

func printMyArrayWithSize(array [5]int) {
	for _, value := range array {
		fmt.Print(value, " ")
	}
}

func printMyArrayWithDynamicSize(array []int) {
	for _, value := range array {
		fmt.Println(value)
	}
}

func shortcoming1() {
	myArray := [5]int{1, 2, 3, 4, 5} // its type is not []int, rather it is [5]int
	fmt.Println(myArray)
	//printMyArray(myArray) // it does not work
	printMyArrayWithSize(myArray) // this works
	//printMyArrayWithDynamicSize(myArray) // needs works
}

func printAddresses(array [5]int) {
	fmt.Printf("Address of passed value : %p\n", &array)
}

func changeArray(array [5]int) {
	array[2] = -1
	printMyArrayWithSize(array)
}

func shortcoming2() {
	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Address of original array : %p\n", &myArray)
	printAddresses(myArray)
	fmt.Printf("confirming Address of original array does not change : %p\n", &myArray)

	fmt.Println("its apparent that chang in array after passing dont reflect in original")
	fmt.Println("original array : ")
	printMyArrayWithSize(myArray)
	fmt.Println()
	fmt.Println("after change, the changed array : ")
	changeArray(myArray) // it also prints the array btw
	fmt.Println()
	fmt.Println("original array :")
	printMyArrayWithSize(myArray)

}

func printTimeWhenArraygotPassed(array [10000000]int) int64 {
	timeInUnix := time.Now().Unix()
	fmt.Println()
	fmt.Println("got inside tha function at time : ", timeInUnix)
	return timeInUnix
}

// Ques: why the time it take to allocte not coming up
//func shortcoming3() {
//	fmt.Println()
//	fmt.Println("Time of start : ", time.Now().Unix())
//	myArray := [10000000]int{};
//	for i := 0; i < 10000000; i++ {
//		myArray[i] = 100
//	}
//	timeStart := time.Now().Unix()
//	fmt.Println()
//	fmt.Println("Time took to create array : ", timeStart)
//	timeEnd := printTimeWhenArraygotPassed(myArray)
//	fmt.Println()
//	fmt.Println("time before method call, took to get inside the func, difference of both : ", timeStart, timeEnd, timeEnd - timeStart)
//}

func shortcoming3() {
	fmt.Println()
	fmt.Println("Time of start : ", time.Now().Unix())
	myArray := make([]int, 10000000) // still does not works!!
	for i := 0; i < 10000000; i++ {
		myArray[i] = 100
	}
	myNewArray := [10000000]int{}
	for i := 0; i < 10000000; i++ {
		myNewArray[i] = myArray[i]
	}
	timeStart := time.Now().Unix()
	fmt.Println()
	fmt.Println("Time took to create array : ", timeStart)
	timeEnd := printTimeWhenArraygotPassed(myNewArray)
	fmt.Println()
	fmt.Println("time before method call, took to get inside the func, difference of both : ", timeStart, timeEnd, timeEnd-timeStart)
}

func shortcomingsOfArray() {
	shortcoming1() // once defined, you canot update its size.
	shortcoming2() // you pass array by values, so when you are passing array, you first create a copy of array then pass (abstracted by golang)
	shortcoming3() //
}

// works fine for array.
func concurrentWritesOnArray() {
	arr := [10]int{}
	for i := 0; i < 10000000; i++ {
		go func() { arr[rand.Intn(10)] = rand.Intn(100000) }()
	}
}

func Start7() {
	myArray := [4]int{1, 2, 3, 4}
	my2dArray := [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	my3dArray := [3][3][2]int{{{1, 2}, {3, 4}, {4, 5}}, {{6, 7}, {8, 9}, {10, 11}}, {{12, 13}, {14, 15}, {16, 17}}}

	fmt.Println("Simple Array")
	for index, value := range myArray {
		fmt.Print(index, ":", value, " ")
	}

	fmt.Println()
	fmt.Println("Simple 2d Array")

	for index, value := range my2dArray {
		fmt.Println(index, ":", value, " ")
		for index2, value2 := range value {
			fmt.Println(index2, ":", value2, " ")
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Simple 3d Array")

	for index, value := range my3dArray {
		fmt.Println(index, ":", value, " ")
		for index2, value2 := range value {
			fmt.Println(index2, ":", value2, " ")
			for index3, value3 := range value2 {
				fmt.Println(index3, ":", value3, " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}

	//havingFunWithIndexIssues()
	//reWritingStackWithNegtiveIndices() // some other time, seems need to work on this!!
	shortcomingsOfArray()

	//concurrent operations on array
	concurrentWritesOnArray()
}
