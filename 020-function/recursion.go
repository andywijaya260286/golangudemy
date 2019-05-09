package main

import(
	"fmt"
)

func main(){
	//factorial
	fmt.Println(4*3*2*1)

	//using recursive func
	x:=factorial(4)
	fmt.Println(x)

	//using for
	k:=factorialloop(4)
	fmt.Println(k)
}

func factorial(n int)int{
	if(n==1){
		return 1
	}
	return n * factorial(n-1)
}

func factorialloop(n int)int{
	k:=1
	for ; n>0; n--{
		k*=n
	}
	return k
}