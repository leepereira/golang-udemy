package main

import "fmt"

type rect struct {
	width  int
	height int
}

func main() {

	myRect := rect{
		width:  12,
		height: 12,
	}

	fmt.Println(myRect.area())

}

// this is a receiver, a parameter of type rect and we named it r
// we creating a area method on the rect struct
func (r rect) area() int {
	return r.width * r.height
}
