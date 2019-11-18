package main

import (
	"fmt"
)

func main() {
	//var x =[]int{1,2,3,4,5}
	x := []int{1, 2, 3, 4, 5}
	for k, v := range x {
		fmt.Println(k, v)
	}
	fmt.Printf("%T\n", x)

}
