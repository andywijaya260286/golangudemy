package main

import(
	"fmt"
	"time"
)

func main(){
	//range pull from channel until its close

	fmt.Println("-----START-----")
	c:=make(chan int)

	//send
	go func(){
		for i:=0;i<5;i++{
			c<-i
			time.Sleep(time.Second*2)
		}
		close(c)
	}()

	//receive
	for x:=range c{
		fmt.Println(x)
	}
	
	fmt.Println("-----END-----")
}