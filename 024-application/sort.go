package main

import(
	"fmt"
	"sort"
)

func main(){
	x := []int{1,7,9,4,6,2,8}
	y := []string{"Budi","Andy","Anjani","Teddy"}

	fmt.Println("before sort =",x)
	sort.Ints(x)
	fmt.Println("after sort =",x)

	fmt.Println("before sort =",y)
	sort.Strings(y)
	fmt.Println("after sort =",y)
}