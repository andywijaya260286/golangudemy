package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string //need to be upper case, otherwise json Marshal not working
	Last  string
	Age   int `umur`
}

func main() {
	p1 := person{
		First: "james",
		Last:  "bond",
		Age:   33,
	}

	p2 := person{
		First: "miss",
		Last:  "moneypeny",
		Age:   29,
	}

	people := []person{
		p1,
		p2,
	}

	fmt.Println(people)

	//func Marshal(v interface{}) ([]byte, error)
	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println("error :", err)
	}

	fmt.Println(string(bs))

	fmt.Println("--------Unmarshal--------")
	//func Unmarshal(data []byte, v interface{}) error
	dataJson := `[{"First":"james","Last":"bond","Age":33},{"First":"miss","Last":"moneypeny","Age":29}]`

	/*var dataJson = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)*/

	dataJsonByte := []byte(dataJson)

	var peopleUnmarshal []person // peopleUnmarshal:=[]person{}

	//err2 := json.Unmarshal(bs, &peopleUnmarshal)
	err = json.Unmarshal(dataJsonByte, &peopleUnmarshal)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("after = %+v \n", peopleUnmarshal) //when printing structs, the plus flag (%+v) adds field names
	fmt.Println("after 2 = ", peopleUnmarshal)
	fmt.Println("after show last data = ", peopleUnmarshal[1:2])

}
