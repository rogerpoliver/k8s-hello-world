package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Handler(response http.ResponseWriter, request *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(response, "<h1>Hello, I'm %s and I'm %s</h1>", name, age)
}

func SecretHandler(response http.ResponseWriter, request *http.Request) {
	password := os.Getenv("password")
	fmt.Fprintf(response, "<h1>The secret is now: %s</h1>", password)
}

func ConfigMap(response http.ResponseWriter, request *http.Request) {
	data, err := os.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Fprintf(response, "My family: %s.", string(data))
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/secret", SecretHandler)
	http.ListenAndServe(":8080", nil)
}
