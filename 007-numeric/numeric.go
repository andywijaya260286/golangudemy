package main

import (
	"fmt"
	"runtime"
)

//use this 2 type below for normal condition
var z float64
var b int

var a int8 = -128 //int8, int16, int32, int64, uint8, uint16, uint32, uint64 (unassign)

func main() {

	//https://golang.org/ref/spec#Numeric_types

	x := 46
	y := 41.448495
	z = 3.14
	b = 1234

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
