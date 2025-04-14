package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Структура для хранения данных пользователя, который авторизован
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// Генерация JWT токена
func GenerateToken(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SECRET_KEY")) // Берем секрет из переменной окружения
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("ошибка при генерации токена: %w", err)
	}

	return tokenString, nil
}

// Парсинг и валидация JWT токена
func ParseToken(tokenString string) (int, error) {
	claims := &Claims{}
	secret := []byte(os.Getenv("JWT_SECRET_KEY")) // Берем секрет из переменной окружения

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("невалидный токен")
	}

	return claims.UserID, nil
}
