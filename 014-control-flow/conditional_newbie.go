package main

import (
	"fmt"
)

func main(){
	if x := 42; x <= 42 {
		fmt.Println("42 == 42")
	} else {
		fmt.Println("our value not 42")
	}

	//switch

	switch{
		case false: fmt.Println("no no")
		case 1==1: fmt.Println("yes"); fallthrough
		case 4==2: fmt.Println("no no 2")
		case 2==2: fmt.Println("yes yes")
	}

	fmt.Println("-----------")

	switch{
		case false: fmt.Println("no no")
		case 2==1: fmt.Println("yes"); 
		case 3==2: fmt.Println("no no 2")
		case 4==3: fmt.Println("yes yes")
		default :	fmt.Println("default")
	}

	fmt.Println("-----------")

	switch "test"{
		case "andy": fmt.Println("andy")
		case "test": fmt.Println("test"); 
		default :	fmt.Println("default")
	}

	fmt.Println("-----------")
	x:=66

	switch{
		case x==55,x==44: fmt.Println("55,44")
		case x==66,x==77: fmt.Println("66,77"); 
		default :	fmt.Println("default")
	}

	fmt.Println("-----------")
	y:=77

	switch y{
		case 55,44: fmt.Println("55,44")
		case 99,77: fmt.Println("77,99"); 
		default :	fmt.Println("default")
	}


}

