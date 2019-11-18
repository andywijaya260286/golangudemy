package main

import (
	"fmt"
)

func main() {
	x := foo()
	fmt.Println(x)
	y, z := bar()
	fmt.Println(y)
	fmt.Println(z)
}

func foo() int {
	return 46
}

func bar() (int, string) {
	return 33, "James"
}
