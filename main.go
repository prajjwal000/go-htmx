package main

import (
	"log"
	"net/http"
)

const (
	port = ":6969"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello on home!!"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on " + port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
