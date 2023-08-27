package main

import (
	"log"
	"time"
)

// User is with Capitals since we might need this user type to be made avaialble to another
// package outside of this package
type User struct {
	FirstName  string
	LastName   string
	PhoneNumbr string
	Age        int
	BirthDate  time.Time
}

func main() {
	user1 := User{
		FirstName: "Leon",
		LastName:  "Pereira",
	}

	log.Println(user1.FirstName, user1.LastName, user1.BirthDate)
}
