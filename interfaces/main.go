package main

import (
	"fmt"
	"interfaces/shapes"
)

func main() {
	// create a Square instance s with a side length of 2.
	s := shapes.Square{Side: 2}
	r := shapes.Rectangle{Length: 2, Breadth: 5}

	fmt.Println(s)
	fmt.Println(r)

	fmt.Println(area(s))
	fmt.Println(area(r))

}

func area(a shapes.Area) int {
	return a.Area()
}
