package main

import(
	"fmt"
)

type person struct{
	first string
	last string
}

type secretagent struct{
	person
	ltk bool
}

func (s secretagent) speak() string{
	return fmt.Sprint("Saya secretagent: ",s.first, s.last)
}

func (p person) speak() string{
	return fmt.Sprint("Saya person: ",p.first, p.last)
}

type human interface{
	speak() (n string)
	// or
	// speak() string
}

func bar(h human){
	fmt.Println("Bar > I called Human",h)
	fmt.Printf("type %T\n",h)
}

func main(){
	x:=secretagent{
		person:person{
			first : "James",
			last : "Bond",
		},
		ltk:true,
	}

	y:=person{
		first : "Jokowi",
		last : "Widodo",
	}

	bar(x)
	bar(y)
}