package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSliceMake("a", a) //a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSliceMake("b", b) // b len=0 cap=5 []

	c := b[:2]
	printSliceMake("c", c) // c len=2 cap=5 [0 0]

	d := c[2:5]
	printSliceMake("d", d) // d len=3 cap=3 [0 0 0]
}

func printSliceMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
