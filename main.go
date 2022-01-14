package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my site made in Golang</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting Server at 3000...")
	http.ListenAndServe(":3000", nil)
}
