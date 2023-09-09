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

	p1 := &v1
	log.Println(p1)
	p1.X = 100
	log.Println(v1)

}
