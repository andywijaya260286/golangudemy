package main

import (
	"fmt"
)

func main() {
	switch {
	case true && true:
		fmt.Println("true&&true = ", true && true)
		fallthrough
	case true && false:
		fmt.Println("true&&false = ", true && false)
		fallthrough
	case true || true:
		fmt.Println("true||true = ", true || true)
		fallthrough
	case true || false:
		fmt.Println("true||false = ", true || false)
		fallthrough
	case !true:
		fmt.Println("!true = ", !true)
	}
}
