package repository

import (
	"context"
	"database/sql"
	"testDealls/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(ctx context.Context, user *domain.User, hashedPassword string) error {

	sql := `INSERT INTO users 
	(name, email, username, hashed_password) 
	VALUES (?, ?, ?, ?)`
	_, err := u.db.ExecContext(ctx, sql,
		user.Name,
		user.Email,
		user.Username,
		hashedPassword,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUser(email string) (domain.User, error) {
	var user domain.User
	row := u.db.QueryRow("SELECT id, hashed_password FROM users WHERE email=?", email)
	if err := row.Scan(&user.ID, &user.Password); err == sql.ErrNoRows {
		return user, err
	} else if err != nil {
		return user, err
	}
	return user, nil
}
