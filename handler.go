package main

import (
	"html/template"
	"net/http"
)

type Post struct {
	Body string
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post := Post{Body: "This is an example test body!"}

	tmpl.Execute(w, post)
}
