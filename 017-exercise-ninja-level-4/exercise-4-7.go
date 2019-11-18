package main

import (
	"fmt"
)

func main() {
	a := []string{"James", "Bond", "007"}
	b := []string{"Money", "Penny", "Mrs"}
	x := [][]string{a, b}
	fmt.Println(x)

	for k, v := range x {
		fmt.Println(k, v)
		for k2, v2 := range v {
			fmt.Println(k2, v2)
		}
	}
}
