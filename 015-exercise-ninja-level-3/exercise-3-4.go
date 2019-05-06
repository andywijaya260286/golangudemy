package main

import (
	"fmt"
)

func main() {
	i := 1986
	for {
		if i <= 2019 {
			fmt.Println(i)
			i++
		} else {
			break
		}
	}
}
