package main

import(
	"fmt"
)

type person struct{
	first string
	last string
	age int
}

func (p person) speak(){
	fmt.Println("IAM HOOMANN, name",p.first,p.last,", age ",p.age)
}

func main(){
	p:=person{
		first:"Wiro",
		last:"Sableng",
		age:33,
	}
	p.speak()
}