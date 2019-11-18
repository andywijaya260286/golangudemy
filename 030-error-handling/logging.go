package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//f, err := os.Create("log.txt")
	//if err!=nil{
	//	fmt.Println("err create file :",err)
	//}
	//defer f.Close()
	//log.SetOutput(f)

	defer foo()

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened :",err)
		//log.Println("err happened :",err) //have additional datetime stamp
		//log.Fatalln(err)
		log.Panicln(err)
	}
	defer f2.Close()

	fmt.Println("check log file in the directory")
}

func foo() {
	fmt.Println("fooo yah")
}
