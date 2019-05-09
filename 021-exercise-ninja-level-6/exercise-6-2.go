package main

import(
	"fmt"
)

func main(){
	x:=[]int{1,2,3,4}
	y:=foo(x...)
	fmt.Println(y)
	z:=bar(x)
	fmt.Println(z)
}

func foo(i ...int)int{
	total := 0
	for _,v:=range i{
		total += v
	}
	return total
}

func bar(i []int)int{
	total := 0
	for _,v:=range i{
		total += v
	}
	return total
}