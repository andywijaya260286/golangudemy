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

func (s secretagent) speak(x string){
	fmt.Println("nama saya : ",s.first, s.last, x)
}

func (x person) walk(){
	fmt.Println(x.first, x.last, " Walking")
}

func (x person) smiling() (ss string){
	return fmt.Sprint(x.first, x.last, " Smiling")
}

func main(){
	a := secretagent{
		person: person{
			first: "Budi",
			last: "Santoso",
		},
		ltk : true,
	}

	a.person.walk()

	y:=a.person.smiling()
	fmt.Println("y = ",y)

	a.speak("hoho")

}