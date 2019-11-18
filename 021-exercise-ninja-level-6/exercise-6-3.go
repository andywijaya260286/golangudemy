package main

import (
	"fmt"
)

func main() {
	defer foo()
	fmt.Println("first")
}

func foo() {
	defer func() {
		fmt.Println("bar")
	}()
	fmt.Println("fooo")
}
