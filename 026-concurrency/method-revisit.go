package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	luas() float64
}

func (c *circle) luas() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("Luas = ", s.luas())
}

func main() {
	// functions with a pointer argument must take a pointer:
	//func ScaleFunc(v *Vertex, f float64) {
	//same with
	//func AbsFunc(v Vertex) float64 { //cannot take address or pointer
	// while methods with pointer receivers take either a value or a pointer as the receiver when they are called
	//func (v *Vertex) Scale(f float64) {
	//almost same with, because actually need to write as (*p)
	//func (v Vertex) Abs() float64 {

	//var v Vertex
	//fmt.Println(v.Abs()) // RUN OK
	//p := &v
	//fmt.Println(p.Abs()) // RUN OK
	//fmt.Println((*p).Abs()) // RUN OK

	c := circle{
		radius: 5,
	}

	d := &c
	e := d.luas()
	fmt.Println(e)

	//or

	x := c.luas()
	fmt.Println(x)

	info(&c)
}
