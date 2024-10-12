package utils

import "strings"

// GetFileExt is a function that returns the extension of a file
// based on the file name.
//
// fileName - the name of the file
func GetFileExt(fileName string) string {
	return fileName[strings.LastIndex(fileName, ".")+1:]
}
