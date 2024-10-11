package server

import (
	"main/internal/service"

	"github.com/labstack/echo/v4"
)

// NewServer is a factory function that returns a new instance of the echo.Echo server
func NewServer() *echo.Echo {
	e := echo.New()

	// Endpoints
	e.POST("/api/v1/savefile", service.SaveFileHandler)

	return e
}
