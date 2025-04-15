package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtSecretKeyEnv = "JWT_SECRET_KEY"

type claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func getTokenTTL() time.Duration {
	hoursStr := os.Getenv("TOKEN_TTL_HOURS")
	hours, err := strconv.Atoi(hoursStr)
	if err != nil || hours <= 0 {
		hours = 24
	}
	return time.Duration(hours) * time.Hour
}

func GenerateToken(userID int) (string, error) {
	ttl := getTokenTTL()

	claims := claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv(jwtSecretKeyEnv))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("ошибка при генерации токена: %w", err)
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (int, error) {
	claims := &claims{}
	secret := []byte(os.Getenv(jwtSecretKeyEnv))

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return 0, fmt.Errorf("невалидный токен: %w", err)
	}

	return claims.UserID, nil
}
