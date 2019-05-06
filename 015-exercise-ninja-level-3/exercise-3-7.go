package main

import (
	"fmt"
)

func main() {
	i := 5
	if i == 1 {
		fmt.Println("Value == 1")
	} else if i == 2 {
		fmt.Println("Value == 2")
	} else {
		fmt.Println("Value not found")
	}
}
