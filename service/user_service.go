package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"testDealls/domain"
	"testDealls/utils"
)

type UserService interface {
	SignUp(ctx context.Context, user *domain.User) error
	Login(user *domain.LoginRequest) (token string, err error)
}

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) SignUp(ctx context.Context, user *domain.User) error {
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return errors.New("all fields are required")
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}

	err = s.repo.CreateUser(ctx, user, hashedPassword)
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			switch sqlErr.Number {
			case 1062:
				return errors.New("username or email already exists")
			default:
				errs := fmt.Sprintf("database error: %s", sqlErr.Message)
				return errors.New(errs)
			}
		}
		return errors.New("internal server error")
	}

	return nil
}

func (s *userService) Login(req *domain.LoginRequest) (token string, err error) {
	u, err := s.repo.GetUser(req.Email)
	if err == sql.ErrNoRows {
		return "", errors.New("invalid credentials")
	} else if err != nil {
		return "", errors.New("Error fetching user")
	}

	if !utils.CheckPasswordHash(req.Password, u.Password) {
		return "", errors.New("invalid credentials")
	}
	token, err = utils.GenerateToken(u)
	if err != nil {
		return "", errors.New("failed generate token")

	}

	return token, nil
}
