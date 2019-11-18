package main

import (
	"fmt"
)

func doSomting(x int) int {
	return x * 5
}

func main() {
	//make(T)          channel    unbuffered channel of type T
	//make(T, n)       channel    buffered channel of type T, buffer size n

	c := make(chan int)

	go func() {
		c <- doSomting(5)
	}()

	fmt.Println(<-c)
}
