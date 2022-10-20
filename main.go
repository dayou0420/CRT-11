package main

import (
	"fmt"
	"net/http"

	"example.com/crt-11/bots"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World 👋!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/callback", bots.Greet)
	http.ListenAndServe(":8080", nil)
}
