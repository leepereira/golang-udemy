package main

import "fmt"

func main() {

	m4 := make(map[string]int)

	m4["Leon"] = 47
	fmt.Println(m4["Leon"])

	m4["Martin"] = 45
	fmt.Println(m4)

	delete(m4, "Leon")
	fmt.Println(m4)

	v, ok := m4["Martin"]
	fmt.Println(v, ok)

	v1, ok1 := m4["Leon"]
	fmt.Println(v1, ok1)

}
