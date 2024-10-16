package service

import (
	"log"
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
		log.Fatal(err)

		return c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to read file",
			Data:       nil,
		})
	}

	defer file.Close()

	// Get the file info
	fileInfo, err := file.Stat()

	if err != nil {
		log.Fatal(err)

		return c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get file info",
			Data:       nil,
		})
	}

	// Create a slice of file data
	var fileData []entity.FileData = make([]entity.FileData, 1)

	// Append the file data to the slice
	fileData[0] = entity.FileData{
		FileName:   fileInfo.Name(),
		FileSize:   fileInfo.Size(),
		FileType:   utils.GetFileType(utils.GetFileExt(fileInfo.Name())).Label(),
		UploadTime: fileInfo.ModTime().String(),
		FileURL:    "/api/v1/files/" + name,
	}

	return c.JSON(http.StatusOK, entity.Response{
		StatusCode: http.StatusOK,
		Message:    "File retrieved successfully",
		Data:       fileData})
}
