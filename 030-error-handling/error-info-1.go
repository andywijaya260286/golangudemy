package main

import(
	"fmt"
	"errors"
)

var em = errors.New("a lebih kecil dari 10")

func main() {
	fmt.Printf("%T\n",em)
	x, err := greaterThen10(4)
	if err!=nil{
		fmt.Println("error : ",err)
	}
	fmt.Println(x)
}

func greaterThen10(a int) (int, error){
	if a<10{
		return a, em
	}

	return a, nil
}