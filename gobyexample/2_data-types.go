package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0 = ", 7/3)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a int8 = -1
	var b uint8 = 2
	var c byte = 3
	var d int16 = -4
	var f uint16 = 5
	var g int32 = -6
	var h rune = -7
	var j uint32 = 8
	var k int64 = -9
	var l uint64 = 10
	var m int = 102
	var n uint = 105

	var f float32 = 18
	var g float32 = 4.5
	var d float64 = 0.23
	var pi float64 = 3.14
	var e float64 = 2.7

	var f complex64 = 1 + 2i
	var g complex128 = 4 + 3i

	var isAlive bool = true
	var isEnabled bool = true

	var name string = "Tom Soier"
	var name2 = "Tom Soier"
	name3 := "Tom Soier"

	var (
		name5 = "Tim"
		age   = 26
	)
	var name6, age2 = "Tom", 25
	name7, age3 := "TOm", 32

}
