package main

import(
	"fmt"
	"sort"
)

type person struct{
	First string
	Age int
}

func (p person) String() string{
	return fmt.Sprintf("%s : %d", p.First, p.Age)
}

type ByAge []person

func (a ByAge) Len() int{ 
	return len(a) 
}

func (a ByAge) Swap(i, j int){ 
	a[i], a[j] = a[j], a[i] 
}

func (a ByAge) Less(i, j int) bool{ 
	return a[i].Age < a[j].Age 
}

//----------------

type ByName []person

func (a ByName) Len() int{ 
	return len(a) 
}

func (a ByName) Swap(i, j int){ 
	a[i], a[j] = a[j], a[i] 
}

func (a ByName) Less(i, j int) bool{ 
	return a[i].First < a[j].First 
}

func main(){
	a := person{
		First:"Rinjani",
		Age:555,
	}

	b := person{
		First:"Bromo",
		Age:111,
	}

	c := person{
		First:"Semeru",
		Age:444,
	}

	x:='A'
	y:='Z'

	fmt.Println(x<y)
	fmt.Println(y<x)

	persons := []person{a,b,c}
	fmt.Println(persons)
	fmt.Println("length = ",ByAge(persons).Len())
	sort.Sort(ByAge(persons))//func Sort(data Interface)
	fmt.Println("order by age = ",persons)

	sort.Sort(ByName(persons))
	fmt.Println("order by name = ",persons)
}