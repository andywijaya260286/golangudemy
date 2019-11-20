package main

import(
	"fmt"
) 

type person struct{
	first string
	last string
}

func main() {
	p1 := person{
		first : "Andy",
		last : "Wijaya",
	}

	p2 := person{
		first : "Teddy",
		last : "Wijaya",
	}

	fmt.Println(p1)
	fmt.Println(p2.first, p2.last)
}