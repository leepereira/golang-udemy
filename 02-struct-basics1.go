package main

import (
	"fmt"
	"log"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	v1 := Vertex{3, 4}
	v2 := Vertex{4, 5}
	log.Println(v1)
	log.Println(v2)
	v1.X = 400
	log.Println(v1.X)

	p1 := &v1 // Create a pointer p1 that points to v1.
	log.Println(p1)
	p1.X = 100      //reassigned the value of v1.X through the pointer p1
	log.Println(v1) //{100 4}

}
