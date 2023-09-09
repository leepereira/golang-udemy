package main

import "log"

func main() {
	// create a Square instance s with a side length of 2.
	var square1 Square = Square{side: 2}
	var rectangle1 Rectangle = Rectangle{length: 3, breadth: 4}

	var square2 Square = Square{side: 2}
	var rectangle2 Rectangle = Rectangle{length: 3, breadth: 4}

	log.Println(areaSquare(square1))
	log.Println(areaRectangle(rectangle1))

	log.Println(areaSquare(square2))
	log.Println(areaRectangle(rectangle2))

}

type Square struct {
	side int
}

type Rectangle struct {
	length  int
	breadth int
}

// s is an instance of the Square struct
//
//	s.side represents the length of one side of the square.
func areaSquare(s Square) int {
	return s.side * s.side
}

func areaRectangle(r Rectangle) int {
	return r.length * r.breadth
}
