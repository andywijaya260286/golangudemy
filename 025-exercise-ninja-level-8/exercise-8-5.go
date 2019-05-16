package main

import(
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type orderAge []user

func(o orderAge) Len() int{
	return len(o)
}
func(o orderAge) Less(i, j int) bool{
	return  o[i].Age < o[j].Age
}
func(o orderAge) Swap(i, j int){
	o[i],o[j] = o[j],o[i]
}

type orderLast []user

func(o orderLast) Len() int{
	return len(o)
}
func(o orderLast) Less(i, j int) bool{
	return  o[i].Last < o[j].Last
}
func(o orderLast) Swap(i, j int){
	o[i],o[j] = o[j],o[i]
}

func main(){
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("before = ",users)

	//sort user by age
	fmt.Println("after sort by age -------- ")
	sort.Sort(orderAge(users))

	for _,v:=range users{
		fmt.Println(v.First,v.Last,v.Age)
		sort.Strings(v.Sayings)
		for _,v2:=range v.Sayings{
			fmt.Println(v2)
		}
		fmt.Println("----")
	}

	//sort user by age
	fmt.Println("after sort by last -------- ")
	sort.Sort(orderLast(users))

	for _,v:=range users{
		fmt.Println(v.First,v.Last,v.Age)
		sort.Strings(v.Sayings)
		for _,v2:=range v.Sayings{
			fmt.Println(v2)
		}
		fmt.Println("----")
	}


}