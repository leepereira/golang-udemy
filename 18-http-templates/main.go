package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_, _ = fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// _ means that if there is a error, just trash that. I dont care about the error
	_ = http.ListenAndServe(portNumber, nil)
}
