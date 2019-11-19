package main

import(
	"fmt"
)

const(
	_ = iota
	kb2 = 1 << (iota * 10)	//iota 1
	mb2 					//iota 2
	gb2						//iota 3
)

func main(){
	//kb := 1024		
	//mb := kb*1024
	//gb := mb*1024

	//fmt.Printf("KB decimal = %d,\tbinary = %b \n", kb, kb)
	//fmt.Printf("KB decimal = %d,\tbinary = %b \n", mb, mb)
	//fmt.Printf("KB decimal = %d,binary = %b \n", gb, gb)

	fmt.Println("kb2 = ",kb2)
	fmt.Println("mb2 = ",mb2)
	fmt.Println("gb2 = ",gb2)
	fmt.Printf("KB decimal = %d,\tbinary = %b \n", kb2, kb2)
	fmt.Printf("KB decimal = %d,\tbinary = %b \n", mb2, mb2)
	fmt.Printf("KB decimal = %d,\tbinary = %b \n", gb2, gb2)


}