package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("Hello Hi")
	}

	g := func(s string) {
		fmt.Println("Hello Hi ", s)
	}

	f()
	g("Symphony")
}
