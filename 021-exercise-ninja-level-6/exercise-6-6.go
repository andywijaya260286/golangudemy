package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Änonymous func")
	}()

	func(i int) {
		fmt.Println("Änonymous func , param =", i)
	}(46)
}
