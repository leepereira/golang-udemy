package main

import "log"

func main() {
	var firstName string
	var age int
	var isMarried bool

	lastName := "Pereira"
	herAge := 15
	isSingle := false

	middleName, hisAge, isBrave := "Leon", 18, true

	log.Println(firstName, age, isMarried)
	log.Println(lastName, herAge, isSingle)

	log.Printf("%T %T %T", firstName, age, isMarried) //string int bool

	log.Println(middleName, hisAge, isBrave)

}
