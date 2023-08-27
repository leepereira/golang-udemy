package main

import "log"

type Bird struct {
	Name string
	Age  int
}

func main() {

	pereiraPets := make(map[string]Bird)

	pet1 := Bird{
		Name: "Luna",
		Age:  1,
	}

	pereiraPets["pet1"] = pet1

	pet2 := Bird{
		Name: "Rio",
		Age:  1,
	}

	pereiraPets["pet2"] = pet2

	log.Println("First Pet Name", pet1.Name)
	log.Println("Second Pet Name", pet2.Name)

}
