package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("leon", "grinal")
	fmt.Println(a, b)
}
