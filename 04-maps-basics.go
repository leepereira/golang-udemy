package main

import "fmt"

type Vertex2 struct {
	Lat  float32
	Long float32
}

var m map[string]Vertex2

func main() {
	m = make(map[string]Vertex2)
	m["Pereira Home"] = Vertex2{40.53, 54.42}
	fmt.Println(m["Pereira Home"])
}
