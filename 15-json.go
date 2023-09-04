package main

import (
	"encoding/json"
	"log"
)

type Family struct {
	ID        int    `json:"id"`
	FirstName string `json:"name"`
	Age       int    `json:"age"`
}

func main() {

	myFamily := `
	
	[
  {
    "id": 1,
    "name": "Leon",
    "age": 47
  },
  {
    "id": 2,
    "name": "Grinal",
    "age": 45
  },
  {
    "id": 3,
    "name": "Chelsea",
    "age": 16
  },
  {
    "id": 4,
    "name": "Hayden",
    "age": 13
  }  
]
`
	// We are going to unmarshal the above json into a struct
	// we are going to create a slice since what are getting above may be more then an entry and
	// we would want to record it in that way in this slice called Person
	var unmarshalled []Family
	err := json.Unmarshal([]byte(myFamily), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmarshalled:  %v", unmarshalled)

	// you cannot print a Person struct.
	// you have to make an instance of the struct
	// give it values and then print the values of the keys in the struct

	// The line var myExtendedFamily []Family declares a variable named myExtendedFamily as a slice of Family structs.
	// This means you're creating an uninitialized slice that can hold elements of type Family
	// this is brandnew and is not related to the above
	var myExtendedFamily []Family

	//first i need to create an instance of the type Family
	myExtendedFamily = append(myExtendedFamily, Family{ID: 5, FirstName: "Glen", Age: 35})
	myExtendedFamily = append(myExtendedFamily, Family{ID: 6, FirstName: "Olivia", Age: 32})

	log.Println(myExtendedFamily) //[{5 Glen 35} {6 Olivia 32}]

	myExtendedFamilyJson, err := json.MarshalIndent(myExtendedFamily, "", "  ")
	if err != nil {
		log.Println("Error Marshalling", err)
	}

	log.Println(string(myExtendedFamilyJson))

}
