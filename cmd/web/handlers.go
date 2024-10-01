package main

import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/base.tmpl.html",
	}

	tm, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	err = tm.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func view(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Life will shown dear %d...", id)

}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Try again in next life", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create your destiny"))
}
