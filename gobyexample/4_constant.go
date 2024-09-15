package main 

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 5000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	const pi float64 = 3.1415
	// const sd float64 nizya

	const (
		pi2 float64 = 3.1452
		e float64 = 2.6212
	)

	const ew, w = 3.12, 2.3

	const (
		a2 = 1 // 1
		b // 1
		cr5 // 1
		d24 = 3 // 3
		f // 3
	)
}
