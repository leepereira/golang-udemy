package main

import "log"

func main() {
	var leonarray = [4]string{"l", "g", "c"}
	log.Println(leonarray)
	leonarray[3] = "h"
	log.Println(leonarray)

}
