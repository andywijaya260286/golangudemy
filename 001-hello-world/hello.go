package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello Go!!")
	fmt.Println(n)
	fmt.Println(err)

	// Short declaration operator
	// Declare a variable and Assign a value 
	a := 42
	fmt.Println("nilai a = ", a)
	a = 45 + 1
	fmt.Println("nilai a update = ", a)

	n2, _ := fmt.Println("Hello Go 2nd!!")
	fmt.Println(n2)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("i = ", i)
		}
	}

}
