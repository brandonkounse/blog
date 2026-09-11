package main

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/brandonkounse/blog/db"
)

func Home(store *db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		post, err := store.GetNewestPost()
		if err != nil {
			if err == sql.ErrNoRows {
				post = db.Post{
					Content:  "No posts found!",
					Category: "General",
				}
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		tmpl.Execute(w, post)
	}
}

func AdminNew(store *db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl, err := template.ParseFiles("admin.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, nil)
			return
		}

		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			title := r.FormValue("title")
			content := r.FormValue("content")
			category := r.FormValue("category")

			err = store.InsertPost(title, content, category)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}
