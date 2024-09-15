package main

import "math"

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func TaskTwentyFour() float64 {
	a := NewPoint(5, 6)
	b := NewPoint(3, 4)

	dist := math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))

	return dist
}
