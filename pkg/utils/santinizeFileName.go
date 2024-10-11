package utils

import "strings"

// SantinizeFileName is a function that santinizes the file name
//
// fileName - the name of the file
func SantinizeFileName(fileName string) string {
	// Replace all spaces with underscores
	fileName = strings.ReplaceAll(fileName, " ", "_")

	return fileName
}
