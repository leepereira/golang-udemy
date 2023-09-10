package main

import "fmt"

type Vertex3 struct {
	Lat1  float32
	Long1 float32
}

var m1 = map[string]Vertex3{
	"Pereira home": Vertex3{40.53, 54.42},
	"Google Home":  Vertex3{40.33, 54.33},
}

func main() {
	fmt.Println(m1)
}
