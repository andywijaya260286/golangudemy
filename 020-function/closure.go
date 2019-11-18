package main

import (
	"fmt"
)

func main() {
	//code block
	{
		x := 678
		fmt.Println(x)
	}

	x := incremental()
	fmt.Println(x)
	fmt.Println(x)

	y := funcincremental()
	fmt.Println("after")
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())

}

func incremental() int {
	var x int
	x++
	return x
}

func funcincremental() func() int {
	var x int
	fmt.Println("inside funcincre = ", x)
	return func() int {
		fmt.Println("inside inside funcincre = ", x)
		x++
		return x
	}
}
