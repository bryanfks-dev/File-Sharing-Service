package service

import (
	"main/core/constant"
	"main/domain/entity"
	"main/pkg/utils"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// GetFileHandler is a handler function that handles the get file endpoint
// It reads the file from the server storage and returns the file as an attachment
func GetFileHandler(c echo.Context) error {
	// Get the name of the file from the URL
	name := c.Param("name")

	// Check if the name is empty
	if name == "" {
		return c.JSON(http.StatusBadRequest, entity.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "File name is required",
			Data:       nil,
		})
	}

	// Open the file
	file, err := os.Open("./" + constant.FILES_DIR + name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to read file",
			Data:       nil,
		})
	}

	defer file.Close()

	return c.Attachment(file.Name(), utils.RemoveDirName(file.Name()))
}
