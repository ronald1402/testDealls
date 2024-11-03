package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"learn/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Setup
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	user := &domain.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Username: "johndoe",
	}

	hashedPassword := "hashed_password_example"

	// Expect the SQL execution
	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Name, user.Email, user.Username, hashedPassword).
		WillReturnResult(sqlmock.NewResult(1, 1)) // LastInsertId=1, RowsAffected=1

	// Run test
	err = repo.CreateUser(context.Background(), user, hashedPassword)

	// Assert
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUser(t *testing.T) {
	// Setup
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	email := "johndoe@example.com"
	expectedUser := domain.User{
		ID:       1,
		Password: "hashed_password_example",
	}

	// Expect the SQL query
	mock.ExpectQuery("SELECT id, hashed_password FROM users WHERE email=?").
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "hashed_password"}).
			AddRow(expectedUser.ID, expectedUser.Password))

	// Run test
	user, err := repo.GetUser(email)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Password, user.Password)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
