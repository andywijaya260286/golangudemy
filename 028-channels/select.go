package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----START-----")
	even := make(chan int) //genap
	odd := make(chan int)
	quit := make(chan int)

	//sent
	go assigningValue(odd, even, quit)

	//receive
	printValue(odd, even, quit)

	fmt.Println("-----END-----")
}

func assigningValue(o, e, q chan<- int) {
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func printValue(o, e, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Genap = ", v)
		case v := <-o:
			fmt.Println("Ganjil = ", v)
		case v := <-q:
			fmt.Println("QUIT = ", v)
			return
		}
	}
}
