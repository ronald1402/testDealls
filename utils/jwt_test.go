package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"learn/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	user := domain.User{
		ID: 12345, // Use int64
	}

	token, err := GenerateToken(user)
	assert.NoError(t, err)

	// Parse the token
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	assert.NoError(t, err)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)

	assert.Equal(t, float64(user.ID), claims["user_id"])
	assert.Equal(t, claims["iat"], float64(time.Now().Unix()))
}
