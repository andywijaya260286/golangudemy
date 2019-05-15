package main

import(
	"fmt"
	"os"
	"io"
)

func main(){
	// func Println(a ...interface{}) (n int, err error) {return Fprintln(os.Stdout, a...)}
	fmt.Println("Hello")
	// func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	fmt.Fprintln(os.Stdout, "Hello") 
	// func WriteString(w Writer, s string) (n int, err error)
	io.WriteString(os.Stdout, "Hello")

	//Everything that implement Write, is also type Writer
	//type Writer interface{
	//	Write(p []byte) (n int, err error)
	//}
}