package main

import(
	"fmt"
)

func main(){
	var a [3]int
	fmt.Println(a)
	a[0]=3
	a[1]=2
	a[2]=1
	fmt.Println(a)
	fmt.Println(len(a))

	// effective go said, use slice instead of array
}