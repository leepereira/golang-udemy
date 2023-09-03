package main

import (
	"log"

	"github.com/leepereira/myprogram/helpers"
)

func main() {
	log.Println("Hello")
	var myVar helpers.SomeType
	myVar.TypeName = "Leon Pereira"
	log.Println(myVar.TypeName)
}
