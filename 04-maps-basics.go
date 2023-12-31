package main

import "fmt"

// first we defined a struct
type Vertex2 struct {
	Lat  float32
	Long float32
}

// declared a map variable named m of type strings
// that associates string keys (locations) with Vertex2 values (coordinates).
var m map[string]Vertex2

func main() {
	// initialized the m map using make to create an empty map
	m = make(map[string]Vertex2)
	m["Pereira Home"] = Vertex2{40.53, 54.42}
	fmt.Println(m["Pereira Home"])
}
