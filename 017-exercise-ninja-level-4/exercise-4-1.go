package main

import (
	"fmt"
)

func main() {
	//x := [5]int{1,2,3,4,5}
	//or can write as
	var x [5]int
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	fmt.Println(x)

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", x)
}
