package main

import (
	"log"
	"net/http"

	"github.com/brandonkounse/blog/db"
	_ "modernc.org/sqlite"
)

func main() {
	// db setup
	store := db.Setup()
	defer store.DB.Close()
	store.CreateTablePosts()

	// mux setup
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)

	fileServer := http.FileServer(http.Dir("assets/"))

	mux.Handle("assets/", http.StripPrefix("assets/", fileServer))

	log.Print("listening on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
