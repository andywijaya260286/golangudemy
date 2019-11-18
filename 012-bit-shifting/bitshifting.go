package main

import (
	"fmt"
)

const (
	_          = iota
	kbiota int = 1 << (iota * 10) // 1 decymal = bit shifting 1*10  = 1 0000 000 000 = 1024 decymal
	mbiota                        //= 1 << (iota*10) // 1 decymal = bit shifting 2*10  = 1 0000 000 000 0000 000 000 = 1048576 decymal
	gbiota                        //= 1 << (iota*10)
)

func main() {
	x := 5
	y := x << 1 // n * 2, x times > 5 * 2, 1 times > 10 or adding 1 zero for binary
	z := x << 2 // n * 2, x times > 5 * 2, 2 times > 20 or adding 2 zero for binary

	c := 10
	a := c >> 1 // n / 2, x times > 10 / 2, 1 times > 5 or removing 1 digit for binary
	b := c >> 2 // n / 2, x times > 10 / 2, 2 times > 2 or removing 2 digit for binary
	d := c >> 3

	fmt.Printf("%d , %b\n", x, x) // 101
	fmt.Printf("%d , %b\n", y, y) // 1010
	fmt.Printf("%d , %b\n", z, z) // 10100

	fmt.Println("------")
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%d , %b\n", c, c) // 1010
	fmt.Printf("%d , %b\n", a, a) // 101
	fmt.Printf("%d , %b\n", b, b) // 10
	fmt.Printf("%d , %b\n", d, d) // 1

	fmt.Println("------")
	//bit 	= 8192
	//byte 	= 1024
	//kilobit 	= 8
	//kilobyte 	= 1
	bt := 8
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d , %b\n", bt, bt) //1000
	fmt.Printf("%d , %b\n", kb, kb) //10000000000
	fmt.Printf("%d , %b\n", mb, mb) //100000000000000000000
	fmt.Printf("%d , %b\n", gb, gb)

	fmt.Println("---Using iota and bit shifting---")
	fmt.Printf("%T, %d , %b\n", kbiota, kbiota, kbiota)
	fmt.Printf("%d , %b\n", mbiota, mbiota)
	fmt.Printf("%d , %b\n", gbiota, gbiota)

}
