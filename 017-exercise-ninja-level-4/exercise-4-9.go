package main

import (
	"fmt"
)

func main() {
	x:= map[string][]string{
		"bond_james":[]string{"Martini","Woman"},
		"moneypenny_miss":{"James Bond","Literatur"},
	}
	x["andy"] = []string{"Hiking","Sepeda"}
	fmt.Println(x)

	for k,v:=range x{
		fmt.Println(k)
		for k2,v2 := range v{
			fmt.Println(k2,v2)
		}
	}
}
