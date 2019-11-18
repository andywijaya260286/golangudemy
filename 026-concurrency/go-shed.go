package main

import (
	"fmt"
	"runtime"
)

func main() {

	//Gosched yields the processor, allowing other goroutines to run.
	//It does not suspend the current goroutine, so execution resumes automatically.

	fmt.Println("Total CPUs = ", runtime.NumCPU())
	fmt.Println("NumGoroutine = ", runtime.NumGoroutine())
	//x:=runtime.GOMAXPROCS(runtime.NumCPU())
	//fmt.Println("GOMAXPROCS = ",x)
	go saySomting("Hello")
	saySomting("World")
}

func saySomting(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
