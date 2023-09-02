package main

import "log"

func main() {
	whoisFamily := "Martin"

	switch whoisFamily {
	case "Leon":
		log.Println("family member is set to Leon")
	case "Grinal":
		log.Println("family member is set to Grinal")
	case "Chelsea":
		log.Println("family member is set to Chelsea")
	case "Hayden":
		log.Println("family member is set to Hayden")

	default:
		log.Println("Not a part of Pereira family")

	}

}
