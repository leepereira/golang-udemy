package main

import (
	"fmt"
	"log"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	log.Println(primes)

	var subslice1 []int = primes[1:4]
	log.Println(subslice1)

	names := []string{"leon", "grinal"}
	log.Println(names)

	// it is a slice containing instances of the struct.
	// this is an anonymous struct
	sliceOfInstancesOfStruct := []struct {
		i int
		b bool
	}{
		{2, false},
		{3, true},
	}
	fmt.Println(sliceOfInstancesOfStruct)

}
