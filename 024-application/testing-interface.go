package main

import (
	"fmt"
)

type Interface interface {
	Length() int
}

func hitungPanjang(i Interface) { //mimic sort.Sort function
	x := i.Length()
	fmt.Println(x)
}

type Car struct {
	brand string
}

type lCar []Car

//Impl dari method Length
func (c lCar) Length() int {
	return len(c)
}

func main() {
	a := Car{
		brand: "Shizuka",
	}
	b := Car{
		brand: "Kawaguchi",
	}
	i := []Car{a, b}
	hitungPanjang(lCar(i))
}
