package chapter01

import "fmt"

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
}