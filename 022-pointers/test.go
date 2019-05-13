package main

import(
	"fmt"
	"math"
)

type circle struct{
	jarijari float64
}

func (c *circle) calculateArea()float64{
	return c.jarijari*c.jarijari*math.Pi
}

type shape interface{
	calculateArea()float64
}

func info(sh shape){
	fmt.Println("Luas = ",sh.calculateArea())
}

func main(){
	lingkaran:=circle{
		jarijari:5,
	}

	// Receiver			Value
	// (t T)			T and *T
	// (t *T)			*T

	// Example
	// (t T)			T and *T
	//info(lingkaran)
	//info(&lingkaran)
	
	// (t *T)			*T
	info(&lingkaran)
}