package main

import (
	"fmt"
)

func main() {
	//make(T)          channel    unbuffered channel of type T
	//make(T, n)       channel    buffered channel of type T, buffer size n

	//they are like runners in a relay race
	//they are synchronized
	//they have to pass/receive the value at the same time
	//the value is passed/received synchronously; at the same time
	//channels allow us to pass values between goroutines

	c := make(chan int)

	//not running
	//fatal error: all goroutines are asleep - deadlock!
	//c<-46

	//using goroutine
	go func() {
		c <- 46
	}()

	fmt.Println("c=", <-c)
}
