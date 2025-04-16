package post_test

import (
	"fmt"
	"testTask/internal/domain"
	"testTask/internal/post/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepositoryInterface(ctrl)

	mockRepo.EXPECT().CreatePost(1, "Post Title", "Post Content", true).Return(1, nil).Times(1)

	postID, err := mockRepo.CreatePost(1, "Post Title", "Post Content", true)

	assert.NoError(t, err)
	assert.Equal(t, 1, postID)
}

func TestCreatePost_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepositoryInterface(ctrl)

	mockRepo.EXPECT().CreatePost(1, "Post Title", "Post Content", true).Return(0, fmt.Errorf("ошибка при создании поста")).Times(1)

	postID, err := mockRepo.CreatePost(1, "Post Title", "Post Content", true)

	assert.Error(t, err)
	assert.Equal(t, 0, postID)
}

func TestGetAllPosts_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepositoryInterface(ctrl)

	expectedPosts := []domain.Post{
		{ID: 1, UserID: 1, Title: "Post 1", Content: "Content 1", CommentsEnabled: true},
		{ID: 2, UserID: 2, Title: "Post 2", Content: "Content 2", CommentsEnabled: false},
	}

	mockRepo.EXPECT().GetAllPosts(10, 0).Return(expectedPosts, nil).Times(1)

	posts, err := mockRepo.GetAllPosts(10, 0)

	assert.NoError(t, err)
	assert.Equal(t, expectedPosts, posts)
}

func TestGetAllPosts_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepositoryInterface(ctrl)

	mockRepo.EXPECT().GetAllPosts(10, 0).Return(nil, fmt.Errorf("ошибка при получении постов")).Times(1)

	posts, err := mockRepo.GetAllPosts(10, 0)

	assert.Error(t, err)
	assert.Nil(t, posts)
}

func TestCreateComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepositoryInterface(ctrl)

	mockRepo.EXPECT().CreateComment(1, "1", "test comment", nil).Return(1, nil).Times(1)

	commentID, err := mockRepo.CreateComment(1, "1", "test comment", nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, commentID)
}

func TestCreateComment_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepositoryInterface(ctrl)

	mockRepo.EXPECT().CreateComment(1, "1", "test comment", nil).Return(0, fmt.Errorf("ошибка при создании комментария")).Times(1)

	commentID, err := mockRepo.CreateComment(1, "1", "test comment", nil)

	assert.Error(t, err)
	assert.Equal(t, 0, commentID)
}

func TestGetCommentsEnabled_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepositoryInterface(ctrl)

	mockRepo.EXPECT().GetCommentsEnabled(1).Return(true, nil).Times(1)

	commentsEnabled, err := mockRepo.GetCommentsEnabled(1)

	assert.NoError(t, err)
	assert.True(t, commentsEnabled)
}

func TestGetCommentsEnabled_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepositoryInterface(ctrl)

	mockRepo.EXPECT().GetCommentsEnabled(1).Return(false, fmt.Errorf("ошибка при получении информации о комментариях")).Times(1)

	commentsEnabled, err := mockRepo.GetCommentsEnabled(1)

	assert.Error(t, err)
	assert.False(t, commentsEnabled)
}
