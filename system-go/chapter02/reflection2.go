package chapter02

import (
	"fmt"
	"reflect"
)

func Start2() {
	type t1 int
	type t2 int

	a := t1(1)
	b := t2(1)
	c := 1

	a_t1 := reflect.ValueOf(&a).Elem()
	b_t2 := reflect.ValueOf(&b).Elem()
	c_t := reflect.ValueOf(&c).Elem()

	fmt.Printf("a type : %s", a_t1.Type())
	fmt.Println()
	fmt.Printf("b type : %s", b_t2.Type())
	fmt.Println()
	fmt.Printf("c type : %s", c_t.Type())
	fmt.Println()
}
