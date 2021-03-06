package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{
		sideLength: 5.5,
	}
	t := triangle{
		height: 3,
		base:   2.1,
	}
	printArea(s)
	printArea(t)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("The area of the shape is:", s.getArea())
}
