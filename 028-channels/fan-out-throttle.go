package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 50; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1 chan int, c2 chan string) {
	var wg sync.WaitGroup
	const goroutines = 5
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(i2 int) {
			i := 0
			for v := range c1 {
				i++
				fmt.Println(i2, v)
				func(v2 int) {
					c2 <- fmt.Sprint(strconv.Itoa(i2) + " job id " + strconv.Itoa(v2) + " ,result =" + strconv.Itoa(timeConsumingWork(v2)))
				}(v)
			}
			fmt.Println(i2, "Total record = ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
