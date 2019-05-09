package main

import(
	"fmt"
)

func main(){
	x:=func(){
		fmt.Println("missing foo")		
	}
	
	x()

	fmt.Printf("%T\n",x)
}