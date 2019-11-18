package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		"bond_james":      []string{"Martini", "Woman"},
		"moneypenny_miss": {"James Bond", "Literatur"},
		"andy":            {"Hiking", "Sepeda"},
	}
	fmt.Println(x)

	delete(x, "andy")

	for k, v := range x {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println(k2, v2)
		}
	}
}
