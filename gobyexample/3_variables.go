package main 

import "fmt"

var fg = "ww"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true 
	fmt.Println(d)

	var e int 
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// var a2, b2, c2 string 

	// var hello string 
	// hello = "hello world"

	var (
		name string
		age int 
	)

	name = "ww"
	age = 21
	fmt.Println(name, age)
}
