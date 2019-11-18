package main

import (
	"fmt"
	"math"
)

type square struct {
	panjang float64
	lebar   float64
}

type circle struct {
	jarijari float64
}

func (c circle) calculateArea() float64 {
	return c.jarijari * c.jarijari * math.Pi
}

func (s square) calculateArea() float64 {
	return s.panjang * s.lebar
}

type shape interface {
	calculateArea() float64
}

func info(sh shape) {
	/*switch sh.(type){
		case circle:
			fmt.Println("luas circle = ", sh.(circle).jarijari*sh.(circle).jarijari*3.14)
		case square:
			fmt.Println("luas square = ", sh.(square).panjang*sh.(square).lebar)
	}*/
	fmt.Println("Luas = ", sh.calculateArea())
}

func main() {
	segiempat := square{
		panjang: 10,
		lebar:   4,
	}

	lingkaran := circle{
		jarijari: 5,
	}

	info(segiempat)
	info(lingkaran)
}
