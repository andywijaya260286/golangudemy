package main

import(
	"fmt"
)

func main(){
	//func (receiver) identifier(parameters) (returns(s)) { ... }

	//func()
	//func(x int) int
	//func(a, _ int, z float32) bool
	//func(a, b int, z float32) (bool)
	//func(prefix string, values ...int)
	//func(a, b int, z float64, opt ...interface{}) (success bool)
	//func(int, int, float64) (float64, *[]int)
	//func(n int) func(p *T)

	defer foo("Xxx") //pending the execution of function until it reached the end

	x:= bar("Budi")
	fmt.Println("a= ",x);

	a, b := foobar("andy","wijaya")
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	// variadic parameter
	v := varidic(2,3,4,5,6,7)
	fmt.Println("sum v = ",v)

	//unfurling a slice
	v2 := varidic([]int{2,3,4,5,6,7}...)
	fmt.Println("sum v2 = ",v2)

	// test
	w := test([]int{1,2,3,4})
	fmt.Println("sum w = ",w)

}

func foo(s string){
	fmt.Println("Hi, ",s)
}

func bar(s string) string{
	return fmt.Sprint("hallo, ",s)
}

func foobar(first, last string)(string, bool){
	a := fmt.Sprint(first, last, ` says "Hello"`)
	b := true
	return a, b
}

//variadic parameters, must be final param
func varidic(i ...int) (int){// can take 0 or more value type int 
	fmt.Println("i, ",i )
	fmt.Printf("i type = %T \n",i)
	sum := 0
	for i,v := range i{
		sum += v
		fmt.Println("Index ke ",i, ", valuenya = ",v, ". Total = ",sum )
	}
	return sum
}

func test(i []int)(int){
	sum := 0
	for i,v := range i{
		sum += v
		fmt.Println("Xndex ke ",i, ", valuenya = ",v, ". Total = ",sum )
	}
	return sum
}