package main

import (
    "net/http"
)

func handler(response http.ResponseWriter, request *http.Request) {
    response.Write([]byte("<h1>Hello, World!</h1>"))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}