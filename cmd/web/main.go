package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct {
	addr string
	staticDir string
}

func main() {
	
	//flags
	var con config
	flag.StringVar(&con.addr,"addr", ":6969", "Network address")
	flag.StringVar(&con.staticDir,"static-dir", "./ui/static", "Static Directory Path")
	flag.Parse()
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(con.staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/life/view", view)
	mux.HandleFunc("/life/create", create)

	log.Println("Starting server on " + con.addr)
	err := http.ListenAndServe(con.addr, mux)
	log.Fatal(err)
}
