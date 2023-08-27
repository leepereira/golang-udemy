package main

import "log"

func main() {

	var leonString string
	leonString = "Green"

	log.Println("leonString is set to", leonString)

	usingPointers(&leonString) // passing the memory address of the leonString variable
	log.Println("after function call, leonString is set to", leonString)

}

func usingPointers(s *string) { // you log the memory address that the pointer s is pointing to.
	log.Println("s is set to ", s)
	newValue := "Red"
	*s = newValue // dereference the pointer s using the * operator and you update the value at the memory location pointed to by the pointer s

}

// Passing large structs by value can be inefficient in terms of memory and performance.
// By using pointers, you can avoid copying the entire struct's data and instead pass a reference to the original data.
