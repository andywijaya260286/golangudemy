package main

import (
	"fmt"
)

func main() {
	//x := type{values} // composite literal
	x := []int{1, 2, 3, 4}
	//slice allows u to group together VALUES of same TYPE
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[0])

	//iterate over slice
	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := range x {
		fmt.Println(i)
	}

	//slicing a slice
	fmt.Println(x[1:4])

	//append to a slice
	a := []int{1, 2, 3}
	fmt.Println(a)
	a = append(a, 4, 5, 6)
	fmt.Println(a)
	//or
	b := []int{8, 88, 888}
	c := []int{9, 99, 999}
	fmt.Println(b)
	fmt.Println(c)
	b = append(b, c...)
	fmt.Println(b)

	//deleting from a slice
	b = append(b[:1], b[4:5]...) // b[:1] == b[0:1]
	fmt.Println(b)

	//make a slice
	fmt.Println("----make-----")
	j := make([]int, 10, 10) // make(type, length, capacity),
	//capacity means: after length reach 10, it will create another 10/capacity size
	fmt.Println(j)
	fmt.Println(len(j))
	fmt.Println(cap(j))

	j[0] = 1
	j[9] = 9
	//cannot add j[10]=10, instead use append
	j = append(j, 3)
	fmt.Println(j)
	fmt.Println(len(j))
	fmt.Println(cap(j))

	j = append(j, 4)
	fmt.Println(j)
	fmt.Println(len(j))
	fmt.Println(cap(j))

	fmt.Println("----multi-dimensional slice-----")
	aw := []string{"andy", "jl riau no 11c", "mazda 6"}
	tw := []string{"teddy", "jl sudirman no 9", "c200"}
	fmt.Println(aw)
	fmt.Println(tw)
	multi := [][]string{aw, tw}
	fmt.Println(multi)
}
