package main

import(
	"fmt"
)

func main(){
	foo("Todd")
	s:= bar("Schmee")
	fmt.Println(s)
	x,y := foobar("Andy", "Wijaya")
	fmt.Println(x)
	fmt.Println(y)

}

func foo(s string){
	fmt.Println("hello ",s)
}

func bar(s string) string{
	return fmt.Sprint("i am ",s)
}

func foobar(first, last string) (string, string){
	return fmt.Sprint(first, last," is my name"), ":D"
}