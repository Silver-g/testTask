package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testTask/pkg/response"
)

func ValidateToken(r *http.Request) (int, error) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf(response.ErrTokenRequired)
	}

	tokenParts := strings.Split(authHeader, "Bearer ")
	if len(tokenParts) != 2 {
		return 0, fmt.Errorf(response.ErrInvalidTokenFormat)
	}

	userID, err := ParseToken(tokenParts[1])
	if err != nil {
		return 0, fmt.Errorf(response.ErrInvalidToken)
	}

	return userID, nil
}
