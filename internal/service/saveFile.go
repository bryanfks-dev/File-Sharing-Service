package service

import (
	"io"
	"log"
	"main/core/constant"
	"main/domain/entity"
	"main/pkg/utils"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// SaveFileHandler is a handler function that handles the save file endpoint
// It saves the uploaded file to the server storage
func SaveFileHandler(c echo.Context) error {
	// Get uploaded file
	file, err := c.FormFile("upload-file")

	if err != nil {
		log.Fatal(err)

		return c.JSON(http.StatusBadRequest, entity.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "No file uploaded",
			Data:       nil,
		})
	}

	// Open the file
	src, err := file.Open()

	if err != nil {
		log.Fatal(err)

		return c.JSON(http.StatusBadRequest, entity.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Error opening the file",
			Data:       nil,
		})
	}

	defer src.Close()

	// Sanitize the file name
	sanitizedFileName := utils.SantinizeFileName(file.Filename)

	// Create a new file
	dest, err := os.Create(constant.FILES_DIR + sanitizedFileName)

	if err != nil {
		log.Fatal(err)

		return c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error creating a new file",
			Data:       nil,
		})
	}

	defer dest.Close()

	// Copy the uploaded file to the new file
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)

		return c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error copying the file into new file",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, entity.Response{
		StatusCode: http.StatusOK,
		Message:    "File uploaded successfully",
		Data:       nil,
	})
}
