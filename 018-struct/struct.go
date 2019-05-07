package main

import (
	"fmt"
)

type person struct{
	first string
	last string
	age int
}

type secretagent struct{
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "Andy",
		last: "Wijaya",
		age: 33,
	}

	p2 := secretagent{
		person: person{
			first: "James",
			last: "Bond",
			age: 31,
		},
		ltk: true,
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.age)

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	fmt.Println(p2.age)
	fmt.Println(p2.ltk)

	//or u can use anonymous struct
	p3 := struct{
		first string
		last string
		age int
	}{
		first : "Teddy",
		last : "Wijaya",
		age : 36,
	}

	fmt.Println(p3)

}
