package main

import (
	"log"
)

func main() {

	//this is a slice that we are iterating on
	familyMembers := []string{"Leon", "Grinal", "Chelsea", "Hayden"}

	//i is the current iteration
	// familyMember is the value
	for i, familyMember := range familyMembers {
		log.Println(i, familyMember)
	}

	// when you dont want to see the iteration number
	for _, familyMember := range familyMembers {
		log.Println(familyMember)
	}

	// create a map and iterate over it
	myPets := make(map[string]string)
	myPets["bird"] = "Luna"
	myPets["bird"] = "Pickles"
	myPets["dog"] = "Toby"

	// will only show the name of the pet
	for _, pet := range myPets {
		log.Println(pet)
	}

	// will show the dog or bird
	for petType, pet := range myPets {
		log.Println(petType, pet)
	}

	// iterate over a string
	var leonString = "Leon used to live in Vasai"

	for i, l := range leonString {
		log.Println(i, "-->", l)
	}

	//iterate over custom objects
	type Family struct {
		FirstName string
		LastName  string
		Age       int
	}

	var rebelloFamily []Family //define an empty slice
	rebelloFamily = append(rebelloFamily, Family{"Neville", "Rebello", 51})
	rebelloFamily = append(rebelloFamily, Family{"Romilda", "Rebello", 47})
	rebelloFamily = append(rebelloFamily, Family{"Helen", "Rebello", 84})

	var pereiraFamily []Family //define an empty slice
	pereiraFamily = append(pereiraFamily, Family{"Leon", "Pereira", 47})
	pereiraFamily = append(pereiraFamily, Family{"Grinal", "Pereira", 45})
	pereiraFamily = append(pereiraFamily, Family{"Chelsea", "Pereira", 16})
	pereiraFamily = append(pereiraFamily, Family{"Hayden", "Pereira", 13})

	for _, l := range rebelloFamily {
		log.Println(l.FirstName, l.LastName, l.Age)
	}

	for _, l := range pereiraFamily {
		log.Println(l.FirstName, l.LastName, l.Age)
	}

}
