package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Store struct {
	DB *sql.DB
}

func Setup() *Store {
	db, err := sql.Open("sqlite", "./db/blog.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Store{DB: db}
}
