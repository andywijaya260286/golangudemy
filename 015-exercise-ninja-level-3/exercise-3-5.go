package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 != 0 { //slighly modified to print only value with >0 reminder
			fmt.Printf("value= %v, dibagi 4 sisa = %d\n", i, i%4)
		}
	}
}
