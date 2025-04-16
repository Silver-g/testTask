package user_test

import (
	"fmt"
	"testTask/internal/domain"
	"testTask/internal/user/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)

	expectedUser := &domain.User{ID: 1, Username: "test_user", Password: "hashed_password"}

	mockRepo.EXPECT().GetUserByID(1).Return(expectedUser, nil).Times(1)

	result, err := mockRepo.GetUserByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
}

func TestGetUserByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)

	mockRepo.EXPECT().GetUserByID(1).Return(nil, fmt.Errorf("ошибка при получении пользователя")).Times(1)

	result, err := mockRepo.GetUserByID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCreateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)

	mockRepo.EXPECT().CreateUser("test_user", "password").Return(1, nil).Times(1)

	userID, err := mockRepo.CreateUser("test_user", "password")

	assert.NoError(t, err)
	assert.Equal(t, 1, userID)
}

func TestCreateUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)

	mockRepo.EXPECT().CreateUser("test_user", "password").Return(0, fmt.Errorf("ошибка при создании пользователя")).Times(1)

	userID, err := mockRepo.CreateUser("test_user", "password")

	assert.Error(t, err)
	assert.Equal(t, 0, userID)
}

func TestGetUserByCredentials_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)

	expectedUser := &domain.User{ID: 1, Username: "test_user", Password: "hashed_password"}

	mockRepo.EXPECT().GetUserByCredentials("test_user", "password").Return(expectedUser, nil).Times(1)

	result, err := mockRepo.GetUserByCredentials("test_user", "password")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
}

func TestGetUserByCredentials_InvalidPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)

	mockRepo.EXPECT().GetUserByCredentials("test_user", "wrong_password").Return(nil, fmt.Errorf("неверный пароль")).Times(1)

	result, err := mockRepo.GetUserByCredentials("test_user", "wrong_password")

	assert.Error(t, err)
	assert.Nil(t, result)
}
