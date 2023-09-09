package main

import "fmt"

// Define an interface named Animal with two methods: Speak() and Move().
type Animal interface {
	Speak() string
	Move() string
}

// Create a concrete type named Dog that implements the Animal interface.
type Dog struct {
	Name string
}

// Implement the Speak() method for Dog.
func (d Dog) Speak() string {
	return "Woof!"
}

// Implement the Move() method for Dog.
func (d Dog) Move() string {
	return "Running on four legs"
}

// Create a concrete type named Bird that implements the Animal interface.
type Bird struct {
	Name string
}

// Implement the Speak() method for Bird.
func (b Bird) Speak() string {
	return "Chirp!"
}

// Implement the Move() method for Bird.
func (b Bird) Move() string {
	return "Flying through the air"
}

func main() {
	// Create instances of Dog and Bird.
	dog := Dog{Name: "Buddy"}
	bird := Bird{Name: "Tweetie"}

	// Create a slice of animals and add the Dog and Bird instances.
	// The statement animals := []Animal{dog, bird} means that you are creating a slice of the Animal interface type
	// and adding instances of Dog and Bird to it.
	animals := []Animal{dog, bird}

	// Loop through the animals and have them speak and move.
	for _, animal := range animals {
		fmt.Printf("%s says '%s' and is %s.\n", animal.(Animal).Speak(), animal.(Animal).Move(), animal.(Animal).Speak())
	}
}
