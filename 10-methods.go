package main

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	X float64
	Y float64
}

// v is the receiver
// Abs is the method
//
//	a method is just a function with a receiver argument.
func (v Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// we have created an instance or the struct Vertex5
	v := Vertex5{3, 4}
	// then called the method Abs on the instance of  that struct
	fmt.Println(v.Abs())
}
