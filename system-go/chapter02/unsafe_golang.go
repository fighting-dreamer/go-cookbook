package chapter02

import (
	"fmt"
	"unsafe"
)

func pointerOperations() {
	p := 10
	p1 := &p
	p2 := (*int32)(unsafe.Pointer(p1))
	p3 := &p1

	fmt.Println("value : ", p)
	fmt.Println("pointer of value : ", p1)
	fmt.Println("value via pointer : ", *p1)
	fmt.Println("unsafe pointer of value using pointer : ", p2)
	fmt.Println("value via unsafe pointer : ", *p2)
	fmt.Println("pointer to pointer : ", p3)
	fmt.Println("getting the pointer value from pointer of pointer : ", *p3)

}

func issueInUnsafePointers() {
	var p int64 = 323232323232
	p1 := &p
	p2 := (*int32)(unsafe.Pointer(p1))
	p3 := &p1

	fmt.Println("value : ", p)
	fmt.Println("pointer of value : ", p1)
	fmt.Println("value via pointer : ", *p1)
	fmt.Println("unsafe pointer of value using pointer : ", p2)
	fmt.Println("value via unsafe pointer (its diff from actual value, coz int32 can't hold int64 value completely): ", *p2)
	fmt.Println("pointer to pointer : ", p3)
	fmt.Println("getting the pointer value from pointer of pointer : ", *p3)

}

func Start3() {
	pointerOperations()
	issueInUnsafePointers()
}
