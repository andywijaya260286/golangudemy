package main

import(
	"fmt"
)

func main() {
	s := "Hello World!"
	fmt.Println(s)
	fmt.Printf("%T\n",s)

	// A string value is a (possibly empty) sequence of bytes
	// byte alias for uint8
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n",bs)

	for i:=0; i< len(s); i++{
		fmt.Print(s[i])
		fmt.Printf(" , %U",s[i])
		fmt.Printf(" , %#U \n",s[i])
	}

	x := "世"
	fmt.Println("value x = ",x)
	fmt.Printf("Type x = %T\n",x)

}