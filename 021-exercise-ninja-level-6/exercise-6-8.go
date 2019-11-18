package main

import (
	"fmt"
)

func main() {
	x := returnfunc()
	fmt.Println(x(46))
}

func returnfunc() func(i int) int {
	return func(i int) int {
		return i
	}
}
