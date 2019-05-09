package main

import(
	"fmt"
)

func main(){
	foo()

	func(){
		fmt.Println("This is anonymous func or self execute func")
	}()

	func(s string){
		fmt.Println("My name is ",s)
	}("Andy")
}

func foo(){
	fmt.Println("This is foo")
}