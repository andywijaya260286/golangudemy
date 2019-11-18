package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	go generateNumber(c)

	for x := range c {
		fmt.Println(x)
	}

	fmt.Println("about to exit")
}

func generateNumber(ch chan<- int) {
	x := 3 //number go routine
	var wg sync.WaitGroup
	wg.Add(x)

	for i := 0; i < x; i++ {
		go func(i2 int) {
			for j := 0; j < 10; j++ {
				fmt.Println(i2, j)
				ch <- j
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(ch)
}
