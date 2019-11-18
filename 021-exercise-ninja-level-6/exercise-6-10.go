package main

import (
	"fmt"
)

func main() {
	x := addingvalue2()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func addingvalue2() func() int {
	var k int
	return func() int {
		k += 2
		return k
	}
}
