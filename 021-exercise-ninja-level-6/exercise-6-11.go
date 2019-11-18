package main

import (
	"fmt"
)

func main() {
	//4*3*2*1
	fmt.Println(factorial(4))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
