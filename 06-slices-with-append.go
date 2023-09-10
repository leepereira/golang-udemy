package main

import "fmt"

func main() {
	var s []int
	printSliceAppend(s) // len=0 cap=0 []

	// append work on empty slices
	s = append(s, 0)
	printSliceAppend(s) //len=1 cap=1 [0]

	// the slice grows as needed
	s = append(s, 1)
	printSliceAppend(s) // len=2 cap=2 [0 1]

	// we can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSliceAppend(s) // len=5 cap=6 [0 1 2 3 4]
}

func printSliceAppend(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
