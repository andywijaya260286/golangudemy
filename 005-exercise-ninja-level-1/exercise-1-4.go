package main

import (
	"fmt"
)

type andywijayatype int

var x andywijayatype

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 49
	fmt.Println(x)
}
