package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	const gs = 100

	wg.Add(gs)

	var counter int64

	for i := 0; i < gs; i++ {
		go func() {
			x := atomic.AddInt64(&counter, 1)
			counter = x
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter = ", counter)
}
