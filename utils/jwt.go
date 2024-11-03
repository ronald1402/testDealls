package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"testDealls/domain"
	"time"

	"github.com/google/uuid"
)

var JWTSecret string

func GenerateToken(user domain.User) (string, error) {
	currentTime := time.Now()

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"iat":     currentTime.Unix(),
		"exp":     currentTime.Add(72 * time.Hour).Unix(),
		"jti":     generateUniqueID(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret))
}

func generateUniqueID() string {
	return uuid.New().String()
}
