package main

import (
	"log"
	"net/http"
)

const (
	port = ":6969"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello on home!!"))
}

func view(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View your life flash before eyes"))
}

func create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Try again in next life", http.StatusMethodNotAllowed)

		return
	}

	w.Write([]byte("Create your destiny"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/life/view", view)
	mux.HandleFunc("/life/create", create)

	log.Println("Starting server on " + port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
