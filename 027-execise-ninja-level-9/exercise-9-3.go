package main

import(
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){

	const gs = 100

	wg.Add(gs)

	counter := 0

	
	for i:=0;i<gs;i++{
		go func(){
			x:=counter
			runtime.Gosched()
			x++
			counter = x
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter = ",counter)
}