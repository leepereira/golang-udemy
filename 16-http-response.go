package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes Written %d", n))
	})

	// _ means that if there is a error, just trash that. I dont care about the error
	_ = http.ListenAndServe(":8080", nil)
}
