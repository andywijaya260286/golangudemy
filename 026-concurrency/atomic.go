package main

import(
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main(){
	fmt.Println("Total CPUs = ",runtime.NumCPU())
	fmt.Println("NumGoroutine init = ",runtime.NumGoroutine())

	var counter int64 

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i:=0;i<gs;i++{
		go func(){
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			//fmt.Println("Counter in for  = ",counter)
			fmt.Println("Counter in atom = ",atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("NumGoroutine = ",runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Counter = ",counter)

	//go run -race atomic.go

}