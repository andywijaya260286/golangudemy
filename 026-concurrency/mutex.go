package main

import(
	"fmt"
	"runtime"
	"sync"
)

func main(){

	//A RWMutex is a reader/writer mutual exclusion lock

	fmt.Println("Total CPUs = ",runtime.NumCPU())
	fmt.Println("NumGoroutine init = ",runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i:=0;i<gs;i++{
		go func(){
			mu.Lock()
			v:=counter
			runtime.Gosched()
			v++
			counter=v

			fmt.Println("Counter in for = ",counter)
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("NumGoroutine = ",runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Counter = ",counter)

	//go run -race race-condition.go

}