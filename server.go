package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(response http.ResponseWriter, request *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(response, "<h1>Hello, I'm %s and I'm %s</h1>", name, age)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
