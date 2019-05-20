package main

import(
	"fmt"
)

func main(){
	//make(T)          channel    unbuffered channel of type T
	//make(T, n)       channel    buffered channel of type T, buffer size n

	c:=make(chan int,2)

	//notrunning 
	//c<-46
	//c<-88
	//c<-99

	//running
	c<-46
	c<-88

	fmt.Println("c1=",<-c)
	fmt.Println("c2=",<-c)
}
