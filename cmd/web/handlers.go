package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
/*

//TODO: Handle no result error separartely
	blogs,err := app.blogmodel.Latest()
	if err != nil {
		app.serverError(w,err)
		return
	}

	for _,v := range blogs {
		app.infoLog.Println(v)
	}
*/
	files := []string{
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
		"./ui/html/base.html",
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

	blog, err := app.blogmodel.Get(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	files := []string{
		"./ui/html/pages/view.html",
		"./ui/html/partials/nav.html",
		"./ui/html/base.html",
	}

	tm, err := template.ParseFiles(files...)
	if err!=nil {
		app.serverError(w,err)
		return
	}

	err = tm.ExecuteTemplate(w, "base", blog)
	if err!=nil{
		app.serverError(w,err)
		return
	}
}

func (app *application) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")

	id, err := app.blogmodel.Insert( title, content)
	if err != nil {
		app.serverError(w,err)
		return
	}

	fmt.Fprintf(w,"Id=%d blog created ",id)
}

func (app *application) creation(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/pages/creation.html",
		"./ui/html/partials/nav.html",
		"./ui/html/base.html",
	}

	tm, err := template.ParseFiles(files...)
	if err!=nil {
		app.serverError(w,err)
		return
	}

	err = tm.ExecuteTemplate(w, "base", nil)
	if err!=nil{
		app.serverError(w,err)
		return
	}
}
