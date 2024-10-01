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
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/life/view", view)
	mux.HandleFunc("/life/create", create)

	log.Println("Starting server on " + port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
