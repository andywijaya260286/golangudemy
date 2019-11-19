package main

import(
	"fmt"
)

func main(){

	for i:=0; i<=3; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("horahore")
		break
	}

	x:=1
	for x<10{
		fmt.Println("x",x)
		x++
	}

	for i:=33; i<=90; i++{
		fmt.Printf("i= %d, hexa = %#U \n",i,i)
	}

	

}