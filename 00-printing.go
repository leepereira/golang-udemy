package main

import (
	"fmt"
	"log"
)

func main() {
	MyAge := 45
	leonsagestring := fmt.Sprintf("I am %d years old", MyAge)
	log.Println(leonsagestring)
}
