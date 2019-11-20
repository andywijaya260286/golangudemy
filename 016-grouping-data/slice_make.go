package main

import(
	"fmt"
)

var a []int = make([]int,5,10)
var b []int

func main(){
	fmt.Println(a)
	fmt.Println(len(a))//5
	fmt.Println(cap(a))//10

	a = append(a, 1,2,3,4,5)
	fmt.Println(a)
	fmt.Println(len(a))//10
	fmt.Println(cap(a))//10

	a = append(a, 6,7)
	fmt.Println(a)
	fmt.Println(len(a))//12
	fmt.Println(cap(a))//20

	fmt.Println(b)
	fmt.Println(len(b))//0
	fmt.Println(cap(b))//0

	// slice multidimention [][]
	x:=[]string{"Andy","Cannondale","CAAD10"}
	y:=[]string{"Teddy","BMC","RC01"}
	z:=[][]string{x,y}

	fmt.Println(z)
}