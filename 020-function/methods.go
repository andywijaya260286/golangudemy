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

type human interface{//everything that can speak is human
	speak()
}

func bar(h human){//only human can access this function
	fmt.Println("I AM HOMANNNN")
}

//func (r receiver) identifier(parameters) (return(s)) { code }

func (s secretagent) speak(){//method
	fmt.Println("I am ", s.first, s.last)
}

func (p person) speak(){//method
	fmt.Println("I am ", p.first, p.last)
}

func main(){
	x:=secretagent{
		person:person{
			first:"James", // "James","Bond"
			last:"Bond",
		},
		ltk:true,
	}

	x2:=secretagent{
		person:person{
			first:"Michael", // "James","Bond"
			last:"Jordan",
		},
		ltk:true,
	}

	p:=person{
		first:"Eric",
		last:"Cantona",
	}

	fmt.Println(x)
	x.speak()
	x2.speak()

	fmt.Println("----------")
	bar(x)
	bar(x2)
	bar(p)
}