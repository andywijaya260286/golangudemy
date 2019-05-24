package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(){
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	
	for x:=range c{
		fmt.Println(x)
	}

	fmt.Println("about to exit")
}