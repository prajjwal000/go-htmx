package main

import (
	"log"
	"net/http"
)

const (
	port = ":6969"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/life/view", view)
	mux.HandleFunc("/life/create", create)

	log.Println("Starting server on " + port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
