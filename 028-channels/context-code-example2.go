package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	x:=  make(chan int)
	go func(c chan <-int){
		c<-5
		}(x)
	fmt.Println("xxx = ",<-x)

	defer cancel() // cancel when we are finished

	for n:=range gen(ctx) {
		fmt.Println(n)
		if n==5{
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	fmt.Println("masuk func gen ")
	dst := make(chan int)
	n:=1
	fmt.Println("sebelum go func")
	go func(){
		for{
			select{
			case <- ctx.Done():
				fmt.Println("masuk case done ",n)
				return 
			case dst<-n://default:
				//dst<-n
				fmt.Println("masuk case dst<-n ",n)
				n++ 
			}

		}
	}()
	fmt.Println("setelah go func")
	return dst
}
