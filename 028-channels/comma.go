package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----START-----")
	even := make(chan int) //genap
	odd := make(chan int)
	quit := make(chan bool)

	//sent
	go assigningValue(odd, even, quit)

	//receive
	printValue(odd, even, quit)

	fmt.Println("-----END-----")
}

func assigningValue(o, e chan<- int, q chan<- bool) {
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func printValue(o, e <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("Genap = ", v)
		case v := <-o:
			fmt.Println("Ganjil = ", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("QUIT if= ", v, ok)
				return
			} else {
				fmt.Println("QUIT else= ", v)
			}

		}
	}
}
