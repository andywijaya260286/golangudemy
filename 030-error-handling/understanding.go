package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{ //return &MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	//type error interface {
	//	Error() string
	//}

	//take a look at interfacepoly.go if dont understand

	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
