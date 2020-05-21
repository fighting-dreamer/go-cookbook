package chapter01

import "fmt"

type Coordinate interface {
}

func findCoordinates(coordinate Coordinate) string {
	return fmt.Sprintf("%v", coordinate)
}

type Point2D struct {
	X int
	Y int
}

func (p Point2D) xaxis() int {
	return p.X
}

func (p Point2D) yaxis() int {
	return p.Y
}

type coordinate int

func (p coordinate) xaxis() int {
	return int(p)
}

func (p coordinate) yaxis() int {
	return 0
}

func Start11() {
	// interfaces are both a type (abstract btw) and a set of methods
	p := Point2D{
		X: 100,
		Y: -20,
	}
	fmt.Println(findCoordinates(p))

	x := coordinate(34)
	fmt.Println(findCoordinates(x))
}
