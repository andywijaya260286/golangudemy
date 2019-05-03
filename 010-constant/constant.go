package main

import (
	"fmt"
)

const a = 42
const(
	b float64 = 88.889
	c = "Testing"
)

func main() {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	fmt.Println("c = ",c)

	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)
}
