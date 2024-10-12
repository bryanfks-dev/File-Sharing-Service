package server

import (
	"main/internal/service"

	"github.com/labstack/echo/v4"
)

// NewServer is a factory function that returns a new instance of the echo.Echo server
func NewServer() *echo.Echo {
	e := echo.New()

	// Endpoints
	e.GET("/api/v1/files", service.ListFilesHandler)
	e.POST("/api/v1/savefile", service.SaveFileHandler)
	e.GET("/api/v1/files/:name", service.GetFileHandler)
	e.GET("/api/v1/files/:name/download", service.DownloadFileHandler)
	e.DELETE("/api/v1/files/:name/delete", service.DeleteFileHandler)

	return e
}
