package user

import (
	"database/sql"
)

type Repository interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
}

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repo{db: db}
}

func (r *repo) CreateUser(user *User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	return r.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
}

func (r *repo) GetUserByUsername(username string) (*User, error) {
	query := "SELECT id, username, password FROM users WHERE username = $1"
	row := r.db.QueryRow(query, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
