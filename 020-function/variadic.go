package main

import (
	"fmt"
)

func main() {
	x := sum("test", 1, 2, 3, 4, 5, 6, 7, 8, 10)
	//x:=sum()
	fmt.Println("Total ALL = ", x)
}

func sum(s string, x ...int) int { //variadic means can take 0 or unlimited values, variadic need to be in the last or final parameter
	// example func sum(x ...int, s string) int{ will not working
	// it should be func sum(s string, x ...int) int{
	fmt.Println(s)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
		fmt.Println("Adding value ", v, ", Total =", sum)
	}

	return sum
}
