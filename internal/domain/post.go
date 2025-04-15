package domain

type PostRequest struct {
	Title           string `json:"title"`
	Content         string `json:"content"`
	CommentsEnabled bool   `json:"comments_enabled"`
}

type Post struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	CommentsEnabled bool   `json:"comments_enabled"`
}

type PostWithUserID struct {
	PostRequest
	UserID int `json:"user_id"`
}
