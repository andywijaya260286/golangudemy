package main

import (
	"errors"
	"fmt"
)

func main() {
	x, err := greaterThen10(4)
	if err != nil {
		fmt.Println("error : ", err)
	}
	fmt.Println(x)
}

func greaterThen10(a int) (int, error) {
	if a < 10 {
		return a, errors.New("a lebih kecil dari 10")
	}

	return a, nil
}
