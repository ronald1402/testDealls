package handler

import (
	"database/sql"
	"net/http"
	"testDealls/commons"

	"github.com/labstack/echo/v4"
)

type HealthCheckHandler struct {
	db          *sql.DB
	IsUnHealthy *bool
}

func NewHealthCheckHandler(e *echo.Echo, db *sql.DB, isUnhealthy *bool) {
	handler := &HealthCheckHandler{
		db:          db,
		IsUnHealthy: isUnhealthy,
	}

	e.GET("/health", handler.Check) // Register the signup route
}

func (h *HealthCheckHandler) Check(c echo.Context) error {
	unHealthy := *h.IsUnHealthy

	if unHealthy {
		return c.JSON(http.StatusServiceUnavailable, commons.Response{Message: "service unavailable"})
	}

	// Check database connection
	if err := h.db.Ping(); err != nil {
		return c.JSON(http.StatusServiceUnavailable, commons.Response{
			Code:    503,
			Message: "Unhealthy",
		})
	}
	return c.JSON(http.StatusOK, commons.Response{
		Code:    200,
		Message: "Ok",
	})
}
