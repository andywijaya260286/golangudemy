package main

import (
	"fmt"
	"log"
)

func main() {
	x, err := greaterThen10(4)
	if err != nil {
		log.Fatalln("error : ", err)
	}
	fmt.Println(x)
}

func greaterThen10(a int) (int, error) {
	if a < 10 {
		return a, fmt.Errorf("%v lebih kecil dari 10", a) //errors.New("a lebih kecil dari 10")
	}

	return a, nil
}
