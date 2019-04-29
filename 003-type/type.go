package main

import (
	"fmt"
)

var z string = `andy said "Assholeeeee"`
var a int = 5

func main() {
	fmt.Println(z)
	fmt.Printf("type = %T\n", z)
	fmt.Printf("base 16, with lower-case letters for a-f hex = %x\n", z)
	fmt.Printf("0x for hex = %#x\n", z)

	fmt.Printf("base 2, binary = %b\n", a)

	printz()
}

func printz() {
	fmt.Println("Nilai z = ", z)
}
