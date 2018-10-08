package main

import "fmt"

// Point is a point in 2D space
type Point struct {
	X, Y int
}

// Circle has a Point and a radius centered at the Point
type Circle struct {
	Point
	Radius int
}

// Wheel has a circle and some spokes
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	fmt.Printf("%#v\n", w)
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v\n", w)
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)
}
