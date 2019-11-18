package main

import (
	"fmt"
)

var a int = 10
var b = a << 1

func main() {
	fmt.Printf("%d, %b, %#x \n", a, a, a)
	fmt.Println("---------")
	fmt.Printf("%d, %b, %#x", b, b, b)
}
