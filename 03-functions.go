package main

import "log"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	log.Println(concat("Leon", " Pereira"))
}
