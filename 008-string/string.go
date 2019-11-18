package main

import (
	"fmt"
)

func main() {
	s := "ABS"
	fmt.Println(s)
	fmt.Printf("%T\n", s)


	// A string value is a (possibly empty) sequence of bytes
	// byte alias for uint8
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i])
		fmt.Printf(" , %U", s[i])
		fmt.Printf(" , %X", s[i])
		fmt.Printf(" , %#U \n", s[i])
	}

	k := string([]byte{228,184,150}) // 3 byte require to create char 世
	fmt.Println("k value = ",k) 

	xx := "世"

	for i,v := range xx {
		fmt.Printf("index %v, xx %X\n", i, v)
		fmt.Printf("index %v, xx %#U\n",i, v)
	}

	x := "世"
	fmt.Println("value x = ", x)
	fmt.Printf("Type x = %T\n", x)

}
