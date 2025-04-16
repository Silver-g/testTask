package post

import "testTask/internal/domain"

type PostRepositoryInterface interface {
	CreatePost(userID int, title, content string, commentsEnabled bool) (int, error)
	GetAllPosts(limit, offset int) ([]domain.Post, error)
	CreateComment(userID int, postID, content string, parentID *int) (int, error)
	GetCommentsEnabled(postID int) (bool, error)
}
