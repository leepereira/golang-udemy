package main

import (
	"log"

	"github.com/leepereira/myprogram/helpers"
)

//argument of intChan, type chan int

const numPool = 10000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {

	// create a channel : a place to send information that will be received in other parts of the program
	// this channel can only hold ints
	intChan := make(chan int)
	defer close(intChan)
	// defer means that : whatever comes after this statement,
	// execute that as soon as the current function is done

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
