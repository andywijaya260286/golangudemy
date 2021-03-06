package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	x := sum(xi...)
	fmt.Println("Total ALL = ", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
		fmt.Println("Adding value ", v, ", Total =", sum)
	}

	return sum
}
