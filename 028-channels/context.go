package main

import(
	"fmt"
	"context"
)

func main(){
	x:=context.Background()

	fmt.Println("Context: ",x)
	fmt.Println("Context err: ",x.Err())
	fmt.Printf("Context tp: %T\n",x)

	fmt.Println("-----------")

	x, cancel := context.WithCancel(x)
	fmt.Println("Context: ",x)
	fmt.Println("Context err: ",x.Err())
	fmt.Printf("Context tp: %T\n",x)
	fmt.Println("Cancel: ",cancel)
	fmt.Printf("Cancel tp: %T\n",cancel)

	fmt.Println("-----------")

	cancel()
	fmt.Println("Context: ",x)
	fmt.Println("Context err: ",x.Err())
	fmt.Printf("Context tp: %T\n",x)

}