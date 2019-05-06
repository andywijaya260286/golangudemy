package main

import(
	"fmt"
)

func main(){

	a:=1
	for a<=3{
		if(a==3){
			break
		}
		fmt.Println("value a = ",a)
		a++ // a = a+1
	}

	b:=1
	for{
		if(b<=3){
			fmt.Println("value b = ",b)
		}else{
			break
		}
		b++
	}

	for i:=0; i<=5; i++{
		if i%2==0{
			continue
		}
		fmt.Println(i)
	}

	z:=33
	for{
		if(z>123){
			break
		}
		fmt.Printf("Value dec = %d, hex = %#x, unicode = %#U\n",z,z,z)
		z++
	}
}