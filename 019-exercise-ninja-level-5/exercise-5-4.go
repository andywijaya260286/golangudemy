package main

import (
	"fmt"
)

func main() {
	s := struct{
		first string
		friends map[string]int
		favDrink []string
	}{
		first:"Andy",
		friends:map[string]int{
			"temanbaik":3,
			"teman":4,
		},
		favDrink:[]string{
			"avocado",
			"bearbrand",
		},
	}

	fmt.Println(s)

	for k, v:= range s.friends{
		fmt.Println(k,v)
	}
	for k, v:= range s.favDrink{
		fmt.Println(k,v)
	}
	
}
