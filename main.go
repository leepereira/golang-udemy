package main

import "fmt"

func main() {
	fmt.Println("Hello")

	var whatToSay string

	whatToSay = "Hello Leon"
	var i int

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, whatElse := saySomething()
	fmt.Println(whatWasSaid, whatElse)
}

func saySomething() (string, string) {
	return "something", "alsosomething"
}
