package main

import (
	"fmt"
)

type person struct {
	name string
}

func changeMe(p *person) {
	//(*p).name = "Super Saiya Goku"
	//or
	p.name = "Super Saiya GOKU"
	fmt.Println(p)
}

func main() {
	p1 := person{
		name: "Goku",
	}

	fmt.Println("before change me = ", p1)
	fmt.Println(&p1.name)

	changeMe(&p1)
	fmt.Println("after change me = ", p1)
}
