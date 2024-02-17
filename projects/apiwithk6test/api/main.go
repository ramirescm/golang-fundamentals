package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", nil)
}
