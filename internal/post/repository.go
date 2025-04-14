package post

import (
	"database/sql"
	"fmt"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) CreatePost(userID int, title, content string, commentsEnabled bool) (int, error) {
	var postID int
	query := `INSERT INTO posts (user_id, title, content, comments_enabled) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.DB.QueryRow(query, userID, title, content, commentsEnabled).Scan(&postID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при создании поста: %w", err)
	}
	return postID, nil
}
