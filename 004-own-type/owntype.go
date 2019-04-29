package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog = 99

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	//cant assign a = b , because different type hotdog != int
	//but u can use conversion
	//CONVERSION

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
