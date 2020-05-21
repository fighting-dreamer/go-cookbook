package chapter01

import "fmt"

type complex struct {
	x, y int
}

func squareWithPointers(x *int) {
	*x = *x * *x // x gets new value
}

func Start6() {
	x := -2
	fmt.Println(x)
	fmt.Println(&x)
	squareWithPointers(&x)
	fmt.Println(x)

	c := complex{
		x: 10,
		y: 4,
	}
	fmt.Println(c)
	fmt.Println(&c)
}
