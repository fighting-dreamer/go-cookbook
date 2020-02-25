package chapter01

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	X int
	Y int
	Label string
}

func createStructInsideAndUseIt() {
	type MyStruct struct {
		U int
		V int
		W string
	}
	p1 := MyStruct{U: 10, V:20, W:"MyStruct-1"}
	p2 := MyStruct{}
	p2.W = "MyStruct-2"

	s1 := reflect.ValueOf(&p1).Elem()
	fmt.Println(s1)

	p1Type := s1.Type()
	fmt.Println(p1Type)

	for i := 0; i < p1Type.NumField(); i++ {
		f := s1.Field(i)
		fmt.Printf("%d = %s\n", i, p1Type.Field(i).Name)
		fmt.Printf("%s = %v\n", f.Type(), f.Interface())
	}
}

func testingCongruence1(value MyStruct) {
	fmt.Println("%v : ", value)
}

func Start10() {
	// printing value via reflect
	p1 := MyStruct{X: 10, Y:20, Label:"MyStruct-1"}
	p2 := MyStruct{}
	p2.Label = "MyStruct-2"

	s1 := reflect.ValueOf(&p1).Elem()
	fmt.Println(s1)

	p1Type := s1.Type()
	fmt.Println(p1Type)

	for i := 0; i < p1Type.NumField(); i++ {
		f := s1.Field(i)
		fmt.Printf("%d = %s\n", i, p1Type.Field(i).Name)
		fmt.Printf("%s = %v\n", f.Type(), f.Interface())
	}

	createStructInsideAndUseIt()
	type MyStruct struct {
		A int
		B int
		C string
	}

	// Code below will give you error
	a1 := MyStruct{A: 10, B:20, C:"MyStruct-1"}
	a2 := MyStruct{}
	a2.C = "MyStruct-2"

	q1 := reflect.ValueOf(&a1).Elem()
	fmt.Println(q1)

	a1Type := q1.Type()
	fmt.Println(a1Type)

	for i := 0; i < p1Type.NumField(); i++ {
		f := q1.Field(i)
		fmt.Printf("%d = %s\n", i, a1Type.Field(i).Name)
		fmt.Printf("%s = %v\n", f.Type(), f.Interface())
	}

	//testing congruence of 2 structs
	//testingCongruence1(a1) // this fails
	x := func(value MyStruct) {
		fmt.Println("%v : ", value)
	}
	x(a1)

}


