package main

import (
	"fmt"
)

func main() {
	//passing a function as an argument is called callback
	tot := sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(tot)

	fmt.Println("---print out sum even number (2,4,6)----")
	deretAngka := []int{1, 2, 3, 4, 5, 6, 7}

	toteven := sumeven(sum, deretAngka...)
	fmt.Println(toteven)
}

func sum(x ...int) int {
	total := 0
	fmt.Printf("%T\n", x)
	for _, v := range x {
		total += v
	}
	return total
}

func sumeven(f func(a ...int) int, b ...int) int {
	var c []int
	for _, v := range b {
		if v%2 == 0 {
			c = append(c, v)
		}
	}

	fmt.Println("Angka Genap = ", c)
	totalEven := f(c...)
	return totalEven
}
