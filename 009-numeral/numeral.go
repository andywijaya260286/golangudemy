package main

import (
	"fmt"
)

func main() {
	
	// Base10 (0-9)
	// 10000 1000 100  10   1
	// 10^4  10^3 10^2 10^1 10^0
	// 4  	 2	  4	   2	   2 (42422)

	// Base2 (0-1)
	// 64  32  16  8   4   2   1
	// 2^4 2^4 2^4 2^3 2^2 2^1 2^0
	//     1   0   1   0   1   0 (42)
	// 1   0   0   1   0   0   0 (72)  

	// Base16 (0-9 A-F)
	// 65536 4096 256  16	 1
	// 16^4  16^3 16^2 16^1 16^0
	//						1 (1)
	//						2 (2)
	//						8 (8)
	//						9 (9)
	//						A (10)
	//                      F (15)
	//			  		1	0 (16)
	//					4	8 (72)
	//			  3		8   F (911)

	a:="H"
	fmt.Println(a)
	fmt.Printf("%T\n",a)

	b:=[]byte(a)
	fmt.Println("Decimal = ",b) //72
	fmt.Printf("%T\n",b)

	c:=b[0]
	fmt.Println("Decimal = ",c) //72
	fmt.Printf("%T\n",c)

	//binary
	fmt.Printf("binary = %b\n",c) //1001000 = 72 

	//Hexadecimal or base 16
	fmt.Printf("Hexadecimal = %#x\n",c) //48 = 72
}
