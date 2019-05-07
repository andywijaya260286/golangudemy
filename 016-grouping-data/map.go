package main

import(
	"fmt"
)

func main(){
	mapNameAge := map[string]int{
		"Andy":33,
		"Teddy":36,
		"Siska":31,
	}

	fmt.Println(mapNameAge)
	fmt.Println(mapNameAge["Andy"])

	fmt.Println(mapNameAge["Budy"]) //still return zero value for not found key value
	//checking
	//u can write it as
	var z int
	var okok bool
	z,okok=mapNameAge["Teddy"]
	fmt.Println("map,z =",z)
	fmt.Println("map,okok =",okok)
	//or
	v,ok:=mapNameAge["Budy"]  
	fmt.Println("map,v =",v)
	fmt.Println("map,ok =",ok)

	if v,ok:=mapNameAge["Andy"]; ok {
		fmt.Println("found in map,v =",v)
		fmt.Println("found in map,ok =",ok)
	}else{
		fmt.Println("not found in map")
	}

	if _,ok:=mapNameAge["Budy"]; ok {
		fmt.Println("found in map")
	}else{
		fmt.Println("not found in map, ok =",ok)
	}

	//Add element 
	mapNameAge["Pibijana"]=55 //similar with slice and array == a[0]=1
	mapNameAge["Anthory"]=66
	mapNameAge["Dora"]=7

	for k,v := range mapNameAge {
		fmt.Println(k,v)
	}

	//Delete element
	delete(mapNameAge, "Dora")

	fmt.Println(mapNameAge)

}