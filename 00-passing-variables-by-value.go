package main

import "log"

func main() {
	x := 5
	new_x := increment(x)
	log.Println(new_x)

}

func increment(x int) int {
	return x + 1
}
