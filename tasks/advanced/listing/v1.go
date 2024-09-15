package main

import "fmt"

func main() {
	a := [5]int{32, 1, 3, 4, 51}
	var b []int = a[1:4]
	fmt.Println(b)
}
