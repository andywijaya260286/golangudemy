package main

import(
	"fmt"
)

var tt []int = []int{1,11,11}

func main() {
	// x:= type{values} //composite literal
	x:= []int{1,2,3,4,5} // slice allow you to group together values of same type
	fmt.Println(x)

	for key, value := range x{
		fmt.Printf("key  %d,value = %d \n",key,value)
	}
	//or
	for i:=0; i<len(x); i++{
		fmt.Println("x = ",x[i])
	}

	xx:=append(tt, 2,3,4)
	fmt.Println(xx)

	//slicing a slice
	fmt.Println(x[1:4]) // [start from : end before]

	//append
	x=append(x, 77,88,99)
	fmt.Println("after append ",x)

	//append another slice
	y:= []int{111,222}
	z:=append(x, y...)
	fmt.Println("append another slice ",z) // [1 2 3 4 5 77 88 99 111 222]

	//delete 3, 4 from the slice
	z=append(z[:2],z[4:]...)
	fmt.Println("after delete ",z)

	//make

}