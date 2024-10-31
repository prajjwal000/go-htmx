package main

import (
	"database/sql"
	"flag"
//	"fmt"
	"log"
	"net/http"
	"os"
	dbmodel "serrver/internal"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	blogm    *dbmodel.BlogModel
}

func main() {

	var con config
	flag.StringVar(&con.addr, "addr", ":6969", "Network address")
	flag.StringVar(&con.staticDir, "static-dir", "./ui/static", "Static Directory Path")
	flag.Parse()

	infoLog := log.New(os.Stdout, "Info:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "Error:\t", log.LstdFlags|log.Lshortfile)

	dburl := `postgres://monkey:junky@localhost:5432/blogs`
	db, err := openDB(dburl)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		blogm:    &dbmodel.BlogModel{DB: db},
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(con.staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/life/view", app.view)
	mux.HandleFunc("/life/create", app.create)

	srv := &http.Server{
		Addr:     con.addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Println("Starting server on " + con.addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dburl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dburl)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
