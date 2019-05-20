package main

import(
	"fmt"
)

func main(){
	fmt.Println("-----START-----")
	x:=make(chan int)

	go func(){
		x<-46
		close(x)
	}()

	v, ok := <-x
	fmt.Println("v=",v)
	fmt.Println("ok=",ok)

	v, ok = <-x
	fmt.Println("v2=",v)
	fmt.Println("ok2=",ok)

	
	fmt.Println("-----END-----")
}