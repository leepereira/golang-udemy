package main

import "log"

func main() {
	x, _ := getCoordinates(4, 6)
	log.Println(x)

}

func getCoordinates(x int, y int) (int, int) {
	return 3, 4
}
