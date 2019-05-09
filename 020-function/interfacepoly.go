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
//allow us to do polymorphism
//allow a value to be of more then one type
//everything that can speak, are also type human
//keyword identifier type
type human interface{
	speak()
}

func foobar(h human){
	switch h.(type){
	case secretagent :
		fmt.Println("IAM HOMANNN with secret agent title , name", h.(secretagent).first)
	case person : 
		fmt.Println("IAM HOMANNN title , name", h.(person).first)
	}
	
}

func (s secretagent) speak(){
	fmt.Println("My name is ",s.first, s.last)
}

func (p person) speak(){
	fmt.Println("My name is ",p.first, p.last)
}

func main(){
	x:=secretagent{
		person:person{
			first:"Wiro",
			last:"Sableng",
		},
		ltk:true,
	}

	p:=person{
		first:"Jimmy",
		last:"Lima",
	}

	fmt.Println(x)
	x.speak()
	foobar(x)
	foobar(p)
	foobar(x.person)
}