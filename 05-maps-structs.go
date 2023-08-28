package main

import "log"

type Bird struct {
	Name string
	Age  int
}

func main() {

	pereiraPets := make(map[string]Bird) //[string]Bird -->  [key]value

	//creates an instance of Bird
	pet1 := Bird{
		Name: "Luna",
		Age:  1,
	}

	// instance pet1 is assigned to map pereiraPets using the key "pet1".
	pereiraPets["pet1"] = pet1

	pet2 := Bird{
		Name: "Rio",
		Age:  1,
	}

	// instance pet2 is assigned to map pereiraPets using the key "pet2".
	pereiraPets["pet2"] = pet2

	log.Println("First Pet Name", pereiraPets["pet1"].Name)
	log.Println("Second Pet Name", pereiraPets["pet2"].Name)
	log.Println("Second Pet Name", pereiraPets)

}
