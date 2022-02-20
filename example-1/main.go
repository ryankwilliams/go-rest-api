package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Foo struct {
	Foo string
}

// home serves the endpoint for the main page of the
// web server.
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
	fmt.Println("/ endpoint accessed")
}

// foo serves the endpoint for the foo page of the
// web server
func foo(w http.ResponseWriter, r *http.Request) {
	var data []Foo
	data = []Foo{
		{Foo: "bar"},
	}
	json.NewEncoder(w).Encode(data)
	fmt.Println("/foo endpoint accessed")
}

// registerEndpoints registers the endpoints functions
func registerEndpoints() {
	http.HandleFunc("/", home)
	http.HandleFunc("/foo", foo)
}

// startServer starts the web server
func startServer() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// main
func main() {
	registerEndpoints()
	startServer()
}
