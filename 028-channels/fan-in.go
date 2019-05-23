package main

import(
	"fmt"
	"sync"
)

func main(){
	fmt.Println("-----START-------")
	even := make(chan int)
	odd	:= make(chan int)
	fanin := make(chan int)

	go separateOddEvenValue(even, odd)

	go combineTo1Channel(even, odd, fanin)

	for x:=range fanin{
		fmt.Println(x)
	}

	fmt.Println("-----END-------")
}

func separateOddEvenValue(even, odd chan <- int){
	fmt.Println("-----Entering separateOddEvenValue-------")
	for i:=0; i<50; i++{
		if i%2==0{
			even<-i
		}else{
			odd<-i
		}
	}
	close(even)
	close(odd)
}

func combineTo1Channel(even, odd <- chan int, fanin chan <- int){
	fmt.Println("-----Entering combineTo1Channel-------")
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		for x:=range even{
			fanin <- x
		}
		wg.Done()
	}()

	go func(){
		for x:=range odd{
			fanin <- x
		}
		wg.Done()
	}()
	
	wg.Wait()
	close(fanin)
}