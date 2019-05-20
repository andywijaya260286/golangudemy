package main

import(
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main(){

	const gs = 100

	wg.Add(gs)

	counter := 0

	
	for i:=0;i<gs;i++{
		go func(){
			mu.Lock()
			x:=counter
			x++
			counter = x
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter = ",counter)
}