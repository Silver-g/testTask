package user

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	DB *sql.DB
}

// Создание пользователя в базе данных
func (r *UserRepository) CreateUser(username, password string) (int, error) {
	var userID int
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	err := r.DB.QueryRow(query, username, password).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при создании пользователя: %w", err)
	}
	return userID, nil
}

// Получение пользователя по ID
func (r *UserRepository) GetUserByID(userID int) (*User, error) {
	var user User
	query := `SELECT id, username, password FROM users WHERE id = $1`
	err := r.DB.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении пользователя: %w", err)
	}
	return &user, nil
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}
func (r *UserRepository) GetUserByCredentials(username, password string) (*User, error) {
	var user User
	query := `SELECT id, username, password FROM users WHERE username = $1 AND password = $2`
	err := r.DB.QueryRow(query, username, password).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("неверное имя пользователя или пароль: %w", err)
	}
	return &user, nil
}
