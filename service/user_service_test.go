package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"learn/domain"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *domain.User, hashedPassword string) error {
	args := m.Called(ctx, user, hashedPassword)
	return args.Error(0)
}

func (m *MockUserRepository) GetUser(email string) (domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}

func TestSignUp(t *testing.T) {
	repo := new(MockUserRepository)
	userService := NewUserService(repo)

	tests := []struct {
		name          string
		inputUser     *domain.User
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Successful Signup",
			inputUser: &domain.User{
				Name:     "John Doe",
				Email:    "john@example.com",
				Username: "johndoe",
				Password: "password123",
			},
			mockSetup: func() {
				repo.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Signup with Empty Fields",
			inputUser: &domain.User{
				Name:     "John Doe",
				Email:    "",
				Username: "johndoe",
				Password: "",
			},
			mockSetup:     func() {},
			expectedError: errors.New("all fields are required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := userService.SignUp(context.Background(), tt.inputUser)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
		repo.AssertExpectations(t)
	}
}
