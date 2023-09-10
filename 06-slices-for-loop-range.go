package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// When ranging over a slice, two values are returned for each iteration.
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
