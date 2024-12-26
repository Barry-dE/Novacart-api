package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// Post represents a structure to hold post data
// It includes fields that map to the database columns and JSON keys for API serialization/deserialization
type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAT string   `json:"updated_at"`
}

// PostsStore provides access to post-related operations in the database
type PostsStore struct {
	db *sql.DB // Database connection
}

// Create inserts a new row into the posts table in the database and persists it
// The method returns any error encountered during execution
func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`
	// Executes the query and scans the returned values into the Post struct
	err := s.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, pq.Array(post.Tags)).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAT,
	)

	// If an error occurs during query execution, return the error
	if err != nil {
		return err
	}

	// Return nil if no error occurs, indicating successful insertion
	return nil
}
