package main

import (
	"fmt"
)

type person struct {
	first  string
	last   string
	flavor []string
}

func main() {
	p1 := person{
		first:  "Andy",
		last:   "Wijaya",
		flavor: []string{"Vanilla", "Strawberry"},
	}

	p2 := person{
		first:  "Budi",
		last:   "Cahyono",
		flavor: []string{"Cocholate", "Macha"},
	}

	fmt.Println(p1)
	for _, v := range p1.flavor {
		fmt.Println(v)
	}

	fmt.Println(p2)
	for _, v := range p2.flavor {
		fmt.Println(v)
	}
}
