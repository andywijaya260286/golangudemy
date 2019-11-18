package main

import (
	"fmt"
)

func main() {
	x := make([]string, 8, 8)
	x = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware"}

	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
	//or
	for k, v := range x {
		fmt.Println(k, v)
	}
}
