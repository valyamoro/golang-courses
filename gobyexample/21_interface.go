package main 

import "math"
import "fmt"

type geometry interface {
	area() float64
	perim() float64 
}

type rect2 struct {
	width, height float64 
}

type circle2 struct {
	radius float64 
}

func (r rect2) area() float64 {
	return r.width * r.height 
}

func (r rect2) perim() float64 {
	return 2 * r.width + 2*r.height 
}

func (c circle2) area() float64 {
	return math.Pi * c.radius * c.radius 
}

func (c circle2) perim() float64 {
	return 2 * math.Pi * c.radius 
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect2{width: 3, height: 4}
	c := circle2{radius: 5}

	measure(r)
	measure(c)
}
