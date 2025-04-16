package post

import (
	"database/sql"
	"fmt"
	"testTask/internal/domain"
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

func (r *PostRepository) GetAllPosts(limit, offset int) ([]domain.Post, error) {
	var posts []domain.Post

	query := `
		SELECT id, user_id, title, content, comments_enabled 
		FROM posts
		ORDER BY id ASC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении постов: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post domain.Post
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

func (r *PostRepository) CreateComment(userID int, postID, content string, parentID *int) (int, error) {
	var commentID int
	query := `
		INSERT INTO comments (user_id, post_id, content, parent_id) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id
	`
	err := r.DB.QueryRow(query, userID, postID, content, parentID).Scan(&commentID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при создании комментария: %w", err)
	}
	return commentID, nil
}

func (r *PostRepository) GetCommentsEnabled(postID int) (bool, error) {
	var commentsEnabled bool
	query := `SELECT comments_enabled FROM posts WHERE id = $1`
	err := r.DB.QueryRow(query, postID).Scan(&commentsEnabled)
	if err != nil {
		return false, fmt.Errorf("ошибка при получении информации о комментариях: %w", err)
	}
	return commentsEnabled, nil
}
