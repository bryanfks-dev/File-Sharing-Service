package service

import (
	"log"
	"main/core/constant"
	"main/domain/entity"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// ListFilesHandler is a handler function that handles the list files endpoint
// It reads the files from the files directory and returns the file data
func ListFilesHandler(c echo.Context) error {
	// Read files from the files directory
	files, err := os.ReadDir("./" + constant.FILES_DIR)

	if err != nil {
		log.Fatal(err)

		return c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to read files",
			Data:       nil,
		})
	}

	// Create a slice of fileData
	var fileData []entity.FileData = make([]entity.FileData, 0)

	// Iterate over the files
	for _, file := range files {
		fileInfo, err := file.Info()

		if err != nil {
			log.Default().Println(err)

			continue
		}

		// Append the file data to the fileData slice
		fileData = append(fileData, entity.FileData{
			FileName:   fileInfo.Name(),
			FileSize:   fileInfo.Size(),
			FileType:   file.Type().String(),
			UploadTime: fileInfo.ModTime().String(),
			FileURL:    "/api/v1/files/" + file.Name(),
		})
	}

	return c.JSON(http.StatusOK, entity.Response{
		StatusCode: http.StatusOK,
		Message:    "Files retrieved successfully",
		Data:       fileData,
	})
}
