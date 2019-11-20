package main

import(
	"fmt"
)

func main(){
	m := map[int]string{
		1: "Andy",
		2: "Budy",
		3: "Teddy",
	}

	fmt.Println("value 3= ",m[3])
	fmt.Println("value 4= ",m[4])

	v, ok := m[3]
	fmt.Println("value v= ",v)
	fmt.Println("value ok= ",ok)

	v2, ok2 := m[4]
	fmt.Println("value v2= ",v2)
	fmt.Println("value ok2= ",ok2)

	//or

	if v,ok:=m[4]; ok{
		fmt.Println("value 1 = ",v)
	}else{
		fmt.Println("not found in map")
	}

	// add element to map
	m[5] = "Siska"

	for i,v := range m{
		fmt.Printf("map index = %v, value = %v \n",i,v)
	}

	// delete element in map
	delete(m, 2)

	for i,v := range m{
		fmt.Printf("map after delete, index = %v, value = %v \n",i,v)
	}

	//array
	var aa [3]int
	aa[1] = 12
	fmt.Println("aa = ",aa)

	//slice
	var bb[]int = []int{1,2,3}
	fmt.Println("bb = ",bb)

	//map
	var cc map[string]int = map[string]int{
		"aa" : 1,
		"bb" : 2,
		"cc" : 3,
	}
	fmt.Println("cc=",cc["cc"])

	for i,v:=range cc{
		fmt.Printf("i = %v, v = %v \n",i,v)
	}
}