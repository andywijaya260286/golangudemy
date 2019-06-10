package main

import(
	"fmt"
)

func main() {
	n, err := fmt.Println("Testing")
    if(err!=nil){
        fmt.Println("err = ",err)
    }

    fmt.Println(n)
}