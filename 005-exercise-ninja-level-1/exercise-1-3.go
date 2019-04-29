package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//%T = a Go-syntax representation of the type of the value
	//%t = the word true or false
	//%s = the uninterpreted bytes of the string or slice
	//%d = base 10 / numbering system that usually we used such as 1, 2, 3
	//%v = the value in a default format
	s := fmt.Sprintf("%d, %s, %t", x, y, z)
	fmt.Println(s)

	t := fmt.Sprintf("%v, %v, %v", x, y, z)
	fmt.Println(t)
}
