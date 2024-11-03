package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"learn/domain"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) SignUp(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserService) Login(user *domain.LoginRequest) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func TestSignUp_Success(t *testing.T) {
	e := echo.New()
	user := `{"username": "ronaldjosua34", "password": "Halomoan18!", "email": "ronaldjosua34@gmail.com"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", bytes.NewBufferString(user))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Set content type
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserService)
	handler := &UserHandler{service: mockService}

	// Set expectation for SignUp, note the use of mock.Anything for the context
	mockService.On("SignUp", mock.Anything, mock.MatchedBy(func(u *domain.User) bool {
		return u.Username == "ronaldjosua34" &&
			u.Password == "Halomoan18!" &&
			u.Email == "ronaldjosua34@gmail.com"
	})).Return(nil)

	err := handler.SignUp(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "user created successfully")
}

func TestSignUp_InvalidRequest(t *testing.T) {
	e := echo.New()
	// Test with an invalid user request (e.g., missing fields)
	user := `{"username": "", "password": "", "email": ""}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", bytes.NewBufferString(user))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Set content type
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserService)
	handler := &UserHandler{service: mockService}

	// Set expectation for SignUp to return an error
	mockService.On("SignUp", mock.Anything, mock.Anything).Return(errors.New("all fields are required"))

	err := handler.SignUp(c)

	// Check if an error occurred and assert the response
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "all fields are required")
}

func TestLogin_Success(t *testing.T) {
	e := echo.New()
	loginRequest := `{"email": "test@example.com", "password": "testpass"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBufferString(loginRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserService)
	handler := &UserHandler{service: mockService}

	// Setting up the mock to expect a Login call with the correct user request
	expectedToken := "token123"
	mockService.On("Login", mock.Anything).Return(expectedToken, nil)

	// Call the handler
	err := handler.Login(c)

	// Check if there was no error
	assert.NoError(t, err)

	// Check the response code and body
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	json.NewDecoder(rec.Body).Decode(&response)

	// Verify the token in the response
	assert.Equal(t, expectedToken, response["token"])
}

func TestLogin_InvalidCredentials(t *testing.T) {
	e := echo.New()
	user := `{"email": "test@example.com", "password": "wrongpass"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBufferString(user))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockUserService)
	handler := &UserHandler{service: mockService}

	// Set expectation for Login to return an error for invalid credentials
	mockService.On("Login", &domain.LoginRequest{Email: "test@example.com", Password: "wrongpass"}).Return("", errors.New("invalid credentials"))

	err := handler.Login(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid credentials")
}
