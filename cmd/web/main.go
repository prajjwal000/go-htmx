package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr string
	staticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	
	var con config
	flag.StringVar(&con.addr,"addr", ":6969", "Network address")
	flag.StringVar(&con.staticDir,"static-dir", "./ui/static", "Static Directory Path")
	flag.Parse()

	infoLog := log.New(os.Stdout, "Info:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "Error:\t", log.LstdFlags|log.Lshortfile)

	app := &application {
		infoLog: infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(con.staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/life/view", app.view)
	mux.HandleFunc("/life/create", app.create)

	srv := &http.Server {
		Addr: con.addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Println("Starting server on " + con.addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
