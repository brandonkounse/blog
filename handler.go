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
