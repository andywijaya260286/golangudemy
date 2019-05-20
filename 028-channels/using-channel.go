package main

import(
	"fmt"
)

func main(){
	fmt.Println("-----START-----")
	c:=make(chan int)

	//send
	go foo(c)

	//receive
	bar(c)
	fmt.Println("-----END-----")
}

//send
func foo(c chan<- int){
	fmt.Println("Entering foo")
	c <- 46
}

//receive
func bar(c <-chan int){
	fmt.Println("Entering bar")
	fmt.Println(<-c)
}