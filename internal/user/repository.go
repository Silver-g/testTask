package user

import (
	"database/sql"
	"fmt"
	"testTask/internal/auth"
	"testTask/internal/domain"
)

type UserRepository struct {
	DB *sql.DB
}

//	func (r *UserRepository) CreateUser(username, password string) (int, error) {
//		var userID int
//		query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
//		err := r.DB.QueryRow(query, username, password).Scan(&userID)
//		if err != nil {
//			return 0, fmt.Errorf("ошибка при создании пользователя: %w", err)
//		}
//		return userID, nil
//	}
func (r *UserRepository) CreateUser(username, password string) (int, error) {
	// Хешируем пароль перед сохранением
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return 0, fmt.Errorf("ошибка при хешировании пароля: %w", err)
	}

	var userID int
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	err = r.DB.QueryRow(query, username, hashedPassword).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при создании пользователя: %w", err)
	}
	return userID, nil
}

func (r *UserRepository) GetUserByID(userID int) (*domain.User, error) {
	var user domain.User
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

//	func (r *UserRepository) GetUserByCredentials(username, password string) (*domain.User, error) {
//		var user domain.User
//		query := `SELECT id, username, password FROM users WHERE username = $1 AND password = $2`
//		err := r.DB.QueryRow(query, username, password).Scan(&user.ID, &user.Username, &user.Password)
//		if err != nil {
//			return nil, fmt.Errorf("неверное имя пользователя или пароль: %w", err)
//		}
//		return &user, nil
//	}
func (r *UserRepository) GetUserByCredentials(username, password string) (*domain.User, error) {
	var user domain.User

	// Получаем пользователя по username
	query := `SELECT id, username, password FROM users WHERE username = $1`
	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("пользователь не найден: %w", err)
	}

	// Сравниваем хеш из БД с введённым паролем
	if err := auth.CheckPasswordHash(user.Password, password); err != nil {
		return nil, fmt.Errorf("неверный пароль")
	}

	return &user, nil
}
