package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float32
	Y float32
}

func Abs(v Vertex) float32 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func Scale(*v Vertex, f float32)  {
	v.X = v.X * f
	v.Y = v.Y * f
}


func main()  {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
