package handler

import (
	"net/http"
	"testDealls/commons"
	"testDealls/domain"
	"testDealls/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(e *echo.Echo, service service.UserService) {
	handler := &UserHandler{service}

	api := e.Group("/api/v1")
	api.POST("/signup", handler.SignUp)
	api.POST("/login", handler.Login)
}

func (h *UserHandler) SignUp(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Response{
			Message: "Invalid Request",
		})
	}

	if err := h.service.SignUp(c.Request().Context(), user); err != nil {
		return c.JSON(http.StatusInternalServerError, commons.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, commons.Response{
		Message: "user created successfully",
	})
}

func (h *UserHandler) Login(c echo.Context) error {
	req := new(domain.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Response{
			Message: "Invalid Request",
		})
	}

	token, err := h.service.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, commons.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
