package main

import (
	"fmt"
)

var x bool = true

func main() {
	a := 7
	b := 42

	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a != b)

	x = false
	fmt.Println(x)
}
