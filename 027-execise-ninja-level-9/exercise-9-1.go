package main

import(
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

func main(){
	fmt.Println("Exercise 1")
	fmt.Println("Start CPUs = ",runtime.NumCPU())
	fmt.Println("Start Go Routines = ",runtime.NumGoroutine())
	wg.Add(2)
	go saySomtingx("Andy")
	
	go saySomtingx("Budy")

	fmt.Println("Mid CPUs = ",runtime.NumCPU())
	fmt.Println("Mid Go Routines = ",runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("End CPUs = ",runtime.NumCPU())
	fmt.Println("End Go Routines = ",runtime.NumGoroutine())
	fmt.Println("Exercise 1 END")

}

func saySomtingx(s string){
	for i:=0;i<10;i++{
		fmt.Println("Hi ",s)
	}
	wg.Done()
}