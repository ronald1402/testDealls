package domain

import "context"

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User, hashPassword string) error
	GetUser(email string) (User, error)
}
