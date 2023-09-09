package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type circle1 struct {
	radius float64
}

func (c circle1) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle1) perimeter() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c1 := circle{4.5}
	r1 := rect{5.0, 6.0}
	shapes := []shape{c1, r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
