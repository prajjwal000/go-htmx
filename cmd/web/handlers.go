package main

import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/base.tmpl.html",
	}

	tm, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	
	err = tm.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) view(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Life will shown dear %d...", id)

}

func (app *application) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w , http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create your destiny"))
}
