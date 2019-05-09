package main

import(
	"fmt"
)

func main(){
	//sample return a string
	x:=foo()
	fmt.Println(x)

	//return a func
	y:=returnfunc()
	fmt.Printf("%T\n",y)
	z:=y()
	fmt.Println(z)
	
	//simplified above return a func
	a:=returnfunc()()
	fmt.Println(a)

	//or
	fmt.Println(returnfunc()())

	fmt.Println("------return func with param-------")

	k:=returnfuncparam()
	fmt.Printf("%T\n",k)
	fmt.Println(k(66))
	//or
	fmt.Println(returnfuncparam()(46))
}

func foo()string{
	s := "Hello foo"
	return s
}

func returnfunc() func() int{
	return func() int{
		return 456
	}
}

func returnfuncparam() func(i int) int{
	return func(i int) int{
		return i
	}
}