package main

import(
	"fmt"
)

const(
	a = 1 << iota //0
	b = 1 << iota //1
	c = 3
	d = 1 << iota //3 >> 100
)

func main(){
	x := 6		
	y := x << 1 // << = bit shifting
	z := x << 2

	xx := 6 << 1

	//convert 6 ke Binary = 110 == 6
	//Bit shift 1 = 1100 == 12
	//Bit shift 2 = 11000 == 24

	// 1 = 1
	// 10 = 2
	// 11 = 3
	// 100 = 4
	// 101 = 5
	// 110 = 6 -
	// 111 = 7
	// 1000 = 8
	// 1001 = 9
	// 1010 = 10
	// 1011 = 11
	// 1100 = 12 -
	// 1101 = 13
	// 1110 = 14
	// 1111 = 15
	// 10000 = 16
	// 10001 = 17
	// 10010 = 18
	// 10011 = 19
	// 10100 = 20
	// 10101 = 21
	// 10110 = 22
	// 10111 = 23
	// 11000 = 24 -

	fmt.Println("x = ",x) //5
	fmt.Println("y = ",y) //10
	fmt.Println("z = ",z) //20
	fmt.Println("xx = ",xx) 

	fmt.Printf("x binary= %b\n",x) //5
	fmt.Printf("y binary= %b\n",y) //10
	fmt.Printf("z binary= %b\n",z) //20
	fmt.Printf("xx binary= %b\n",xx) 

	fmt.Println("a = ",a)  //1
	fmt.Println("b = ",b)  //2
	fmt.Println("c = ",c)  //3
	fmt.Println("d = ",d)  //8
}