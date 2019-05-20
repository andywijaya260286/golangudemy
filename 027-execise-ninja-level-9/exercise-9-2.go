package main

import(
	"fmt"
)

type person struct{
	name string
}

func (p *person) speak() string{
	return fmt.Sprintf("Nama saya, %v",p.name)
}

type human interface{
	speak() string
}

func saySomting2(h human){
	fmt.Println("Say : " ,h.speak())
}

func main(){
	a:=person{
		name:"Andy",
	}

	saySomting2(&a)
	//saySomting(a)
}