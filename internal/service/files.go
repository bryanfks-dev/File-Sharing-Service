package service

import (
	"log"
	"main/core/constant"
	"main/domain/entity"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

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
	var fileData []entity.FileData

	// Iterate over the files
	for _, file := range files {
		fileInfo, err := file.Info()

		if err != nil {
			log.Default().Println(err)

			continue
		}

		// Append the file data to the fileData slice
		fileData = append(fileData, entity.FileData{
			FileName:   file.Name(),
			FileSize:   fileInfo.Size(),
			FileType:   file.Type().String(),
			UploadTime: fileInfo.ModTime().String(),
			FileURL:    "/api/v1/getfile/" + file.Name(),
		})
	}

	return c.JSON(http.StatusOK, entity.Response{
		StatusCode: http.StatusOK,
		Message:    "Files retrieved successfully",
		Data:       fileData,
	})
}
