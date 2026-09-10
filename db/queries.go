package db

import "log"

func (s *Store) CreateTablePosts() {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	content TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	category TEXT	
	);	
	`

	_, err := s.DB.Exec(query)
	if err != nil {
		log.Fatalf("failed to create posts table: %v", err)
	}
}
