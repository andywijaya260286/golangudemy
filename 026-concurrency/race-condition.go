package main

import(
	"fmt"
	"runtime"
	"sync"
)

func main(){
	fmt.Println("Total CPUs = ",runtime.NumCPU())
	fmt.Println("NumGoroutine init = ",runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i:=0;i<gs;i++{
		go func(){
			v:=counter
			runtime.Gosched()
			v++
			counter=v

			fmt.Println("Counter in for = ",counter)
			wg.Done()
		}()
		fmt.Println("NumGoroutine = ",runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Counter = ",counter)

	//go run -race race-condition.go

}