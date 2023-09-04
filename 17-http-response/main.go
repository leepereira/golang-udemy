package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home page")

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(4, 3)
	fmt.Fprintf(w, "this is the about page")
	_, _ = fmt.Fprintf(w, fmt.Sprintf("the sum of 4 + 3 is %d", sum))

}

func addValues(x, y int) int {
	return x + y
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_, _ = fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// _ means that if there is a error, just trash that. I dont care about the error
	_ = http.ListenAndServe(portNumber, nil)
}
