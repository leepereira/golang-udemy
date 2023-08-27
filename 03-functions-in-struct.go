// this is also called as a receiver in go

package main

import "log"

type Animal struct {
	Family string
	Mammal bool
	Age    int
}

// this is called a receiver and it ties the function printAnimalFamily to the
// type of my struct Animal
// use case is that the function can be complex and have some busines logic
// that can be a part of this struct

// this shows the use of a pointer receiver with the method printAnimalFamily

// The receiver indicates on which instance the method is called, (in this case : Animal)
// and it can be either a value or a pointer to the type.

// value receiver, it operates on a copy of the instance.

// When a method has a pointer receiver,
//it operates directly on the instance pointed to by the pointer.

// Any changes made within the method will directly affect the original instance outside the method.

func (a *Animal) printAnimalFamily() string {
	return a.Family

}

func (b *Animal) updateAge() int {
	b.Age = 100
	return b.Age
}

func main() {

	var hisPet Animal
	hisPet.Family = "Dog"
	hisPet.Mammal = true
	hisPet.Age = 2

	myPet := Animal{
		Family: "Bird",
		Mammal: false,
		Age:    1,
	}

	log.Println("hisPet is set to ", hisPet.Family)

	// when you call the printAnimalFamily method on the instance of Animal, hisPet in this case
	// you are direcly working with that instance hisPet
	log.Println("hisPet is set to ", hisPet.printAnimalFamily())
	log.Println("myPet is set to ", myPet.Family)
	log.Println("hisPet is set to ", myPet.printAnimalFamily())

	//any change that you make to this instance will update the original instance

	log.Println("hisPet age is now set to ", hisPet.updateAge())

}
