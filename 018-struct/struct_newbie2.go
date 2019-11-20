package main

import(
	"fmt"
) 

type person struct{
	first string
	last string
}

type specialagent struct{
	person
	kill bool
	first string
}

func main() {
	p1 := person{
		first : "Andy",
		last : "Wijaya",
	}

	fmt.Println(p1)

	p2 := specialagent{
		person : p1,
		kill : false,
		first : "pertamax",
	}

	fmt.Println(p2)

	//or
	p3 := specialagent{
		person : person{
			first : "Budi",
			last : "Guguk",
		},
		kill : true,
		first : "keduax",
	}

	fmt.Println(p3)
	fmt.Println(p3.first)
	fmt.Println(p3.person.first)

	//anonymous struct <dont have name>
	p4 := struct{
		nama string
		umur int
	}{
		nama : "susi",
		umur : 22,
	}

	fmt.Println(p4)

}