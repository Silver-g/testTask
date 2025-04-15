package boundary

import (
	"testTask/internal/domain"
)

func MapPostRequestToPostWithUserID(req domain.PostRequest, userID int) domain.PostWithUserID {
	return domain.PostWithUserID{
		PostRequest: req,
		UserID:      userID,
	}
}
func MapPostWithUserIDToPost(postWithUserID domain.PostWithUserID, postID int) domain.Post {
	return domain.Post{
		ID:              postID,
		UserID:          postWithUserID.UserID,
		Title:           postWithUserID.Title,
		Content:         postWithUserID.Content,
		CommentsEnabled: postWithUserID.CommentsEnabled,
	}
}
