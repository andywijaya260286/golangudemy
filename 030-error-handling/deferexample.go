package main

import (
	"fmt"
)

func main() {
	goo()

	fmt.Println("--------------")

	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("--------------")
	z := dool()
	fmt.Println(z)
}

func goo() {
	i := 0
	defer fmt.Println(i) //Last In First Out order
	i++
	fmt.Println(i)
	return
}

func dool() (k int) {
	defer func() { k++ }()
	return 1
}
