package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


// v is a pointer to the original Vertex
// Using a pointer receiver means that you are working directly with the original Vertex instance,
// and any modifications made to v within the method will affect the original Vertex
// So, if you call v.Scale(10) on a Vertex, 
// it will change the original Vertex by scaling its X and Y coordinates.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v.X)
	fmt.Println(v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())
}

// In summary, using a pointer receiver (*Vertex) allows you to modify the actual instance of the Vertex 
// that you call the method on, 
// while a value receiver (Vertex) operates on a copy of the original Vertex.
