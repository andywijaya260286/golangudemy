package main

import (
	"fmt"
)

func main() {

	//1. chan T          // can be used to send and receive values of type T
	//2. chan<- float64  // can only be used to send float64s
	//3. <-chan int      // can only be used to receive ints

	//1
	c := make(chan int, 2)

	c <- 41
	c <- 88

	fmt.Println("c=", <-c)
	fmt.Println("c=", <-c)

	fmt.Printf("type c= %T\n", c)

	//2
	fmt.Println("--------can only be used to send int to a channel---------")
	d := make(chan<- int, 2)

	d <- 41
	d <- 88

	//cant receive as below code
	//fmt.Println("c=",<-d)

	fmt.Printf("type d= %T\n", d)
	fmt.Printf("type cd= %T\n", chan<- int(c))

	//3
	fmt.Println("--------can only be used to receive ints from a channel---------")
	e := make(<-chan int, 2)

	//cant sent as below code
	//e<-41

	fmt.Printf("type e= %T\n", e)
}
