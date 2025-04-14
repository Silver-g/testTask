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

// Получение всех постов
func (r *PostRepository) GetAllPosts() ([]Post, error) {
	var posts []Post
	query := `SELECT id, user_id, title, content, comments_enabled FROM posts`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении постов: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CommentsEnabled); err != nil {
			return nil, fmt.Errorf("ошибка при чтении постов: %w", err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении данных: %w", err)
	}

	return posts, nil
}
