package user

import (
	"testTask/internal/domain"
)

type UserRepositoryInterface interface {
	CreateUser(username, password string) (int, error)
	GetUserByID(userID int) (*domain.User, error)
	GetUserByCredentials(username, password string) (*domain.User, error)
}
