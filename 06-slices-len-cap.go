package main

import "fmt"

func main() {
	s := []int{2, 3, 4, 5, 6, 7, 13}
	printSlice(s) //len=7 cap=7 [2 3 4 5 6 7 13]

	// Slice the slice to give it zero length.
	// This operation slices the slice s to give it a zero length.
	// In other words, it empties the slice.
	s = s[:0]
	printSlice(s) //len=0 cap=7 []

	// Extend its length.
	// This operation extends the length of the slice s to 4
	// by including the first 4 elements from the original slice.
	s = s[:4]
	printSlice(s) //len=4 cap=7 [2 3 4 5]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len=2 cap=5 [4 5]

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
