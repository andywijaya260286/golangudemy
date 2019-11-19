package main

import (
	"fmt"
)

const (
	a = 0
	b = iota
	c
)

const (
	d = iota
	e
	f 
	g = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("-------")
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
