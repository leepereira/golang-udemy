package main

import (
	"fmt"
	"log"
)

func main() {

	var x int
	x = 11

	if x == 10 {
		log.Println("x is 10")
	} else {
		log.Println("x is not 10")
	}

	if x <= 100 || x > 100 {
		log.Println("OR condition tested")
	} else if x == 11 {
		log.Println("x is 11")
	} else {
		fmt.Println("this is a default message")
	}
}
