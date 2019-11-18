package main

import (
	"fmt"
)

func main() {
	x := foo(sum69)

	fmt.Println(x)

	//or can use func as variable
	y := func(i int) int {
		return i + 10
	}
	z := foo(y)
	fmt.Println(z)
}

func sum69(i int) int {
	return i + 4
}

func foo(f func(i int) int) int {
	x := f(46)
	return x
}
