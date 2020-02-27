package chapter01

import "fmt"

func passByValueForSlice(aSlice []int) {
	fmt.Printf("address : %p\n", &aSlice)
}

func makingChanges(aSlice []int) {
	aSlice[2] = -2
}

func printSlice(aSlice []int) {
	for _, v := range aSlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func Start8() {
	//understanding slice, capacity, length
	//slice is array but with dynamic size, if it runs out of capacity, go will assign double 'capacity' of its 'original' more to it
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice, cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	fmt.Println(aSlice, cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -200)
	aSlice = append(aSlice, -300)
	aSlice = append(aSlice, -400)
	fmt.Println(aSlice, cap(aSlice), len(aSlice))
	// you see the capacity got doubled in first append, while the length increased incrementaally with each append

	fmt.Printf("address : %p\n", &aSlice)
	// the address value printed will be different : https://stackoverflow.com/questions/39993688/are-golang-slices-passed-by-value
	// https://golang.org/pkg/reflect/#SliceHeader
	// https://blog.golang.org/go-slices-usage-and-internals
	passByValueForSlice(aSlice)

	fmt.Println()
	printSlice(aSlice)
	makingChanges(aSlice)
	printSlice(aSlice)
}
