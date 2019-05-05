package main

import (
	"fmt"
	"net/http"
)

func userSampleApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to Golang programming language</h1>")
}

func main() {
	http.HandleFunc("/api", userSampleApi)
	http.ListenAndServe(":8000", nil)
}
