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

	mapPerson := map[string]person{}

	mapPerson[p1.last] = p1
	mapPerson["Cahyono"] = p2

	for _, v := range mapPerson {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for _, v2 := range v.flavor {
			fmt.Println(v2)
		}
		fmt.Println("--------")
	}
}
