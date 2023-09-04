package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_, _ = fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// _ means that if there is a error, just trash that. I dont care about the error
	_ = http.ListenAndServe(portNumber, nil)
}
