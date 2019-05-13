package main

import (
	"fmt"
)

func main() {
	a := 46
	fmt.Println(a)
	fmt.Println(&a) // & give you the address

	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",&a)

	//b:=&a 
	var b *int = &a
	fmt.Println(*b)
	fmt.Println(*&a) // * give u the value from the address

	*b = 88
	fmt.Println(a)
	fmt.Println(*b)

	//use pointer when u dont want to pass big chunk of data, u can pass the address

	//*Everything in go pass by value
	fmt.Println("------pass by value example------")
	c:=46
	foo(c)
	fmt.Println(c)

	fmt.Println("------pass by value example with pointers------")
	bar(&c)
	fmt.Println("end after=",c)

	fmt.Println("------pass by value example with pointers 2------")
	foo(*&c)
	fmt.Println("end after=",c)

}

func foo(x int){
	fmt.Println(x)
	x = 17
	fmt.Println(x)
}

func bar(x *int){
	fmt.Println("initial bar=",x)
	fmt.Println("initial bar=",*x)
	*x = 17
	fmt.Println("after bar=",x)
	fmt.Println("after bar=",*x)
}

