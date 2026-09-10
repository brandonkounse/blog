package db

import (
	"log"
	"time"
)

type Post struct {
	Title       string
	Content     string
	DateCreated time.Time
	DateUpdated time.Time
	Category    string
}

func (s *Store) CreateTablePosts() {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME,
	category TEXT	
	);	
	`

	_, err := s.DB.Exec(query)
	if err != nil {
		log.Fatalf("failed to create posts table: %v", err)
	}
}

func (s *Store) GetNewestPost() (Post, error) {
	query := `
	SELECT 
		title,
		content,
		created_at,
		updated_at,
		category
	FROM posts
	ORDER BY created_at DESC
	LIMIT 1;
	`

	var post Post

	err := s.DB.QueryRow(query).Scan(&post.Title, &post.Content, &post.DateCreated, &post.DateUpdated, &post.Category)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}
