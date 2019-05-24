package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring2("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			fmt.Println("random ",msg,i,time.Duration(rand.Intn(1e3)))
			//time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			time.Sleep(time.Second*1)
		}
	}()
	return c
}

func boring2(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			fmt.Println("random ",msg,i,time.Duration(rand.Intn(1e3)))
			//time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			time.Sleep(time.Second*3)
		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/
