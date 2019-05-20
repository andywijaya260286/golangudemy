package main

import(
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func init(){
	fmt.Println("OS = ",runtime.GOOS)
}

func main(){

	//A WaitGroup waits for a collection of goroutines to finish. 
	//The main goroutine calls Add to set the number of goroutines to wait for. 
	//Then each of the goroutines runs and calls Done when finished. 
	//At the same time, Wait can be used to block until all goroutines have finished.

	fmt.Println("Total CPUs = ",runtime.NumCPU())
	fmt.Println("Total Go Routines = ",runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

	fmt.Println("CPU Architecture = ",runtime.GOARCH)
	

	fmt.Println("Total CPUs = ",runtime.NumCPU())
	fmt.Println("Total Go Routines = ",runtime.NumGoroutine())
	

}

func foo(){
	for i:=1; i<10; i++{
		fmt.Println("foo = ",i)
	}
	wg.Done()
}

func bar(){
	for i:=1; i<10; i++{
		fmt.Println("bar = ",i)
	}
	wg.Done()
}