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
	fmt.Println("Total CPUs = ",runtime.NumCPU())
	fmt.Println("Total Go Routines = ",runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	fmt.Println("Arch = ",runtime.GOARCH)
	bar()

	fmt.Println("Total CPUs = ",runtime.NumCPU())
	fmt.Println("Total Go Routines = ",runtime.NumGoroutine())
	wg.Wait()

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
}