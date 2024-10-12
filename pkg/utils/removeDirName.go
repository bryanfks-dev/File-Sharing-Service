package utils

import (
	"main/core/constant"
	"strings"
)

// removeDirName is a function that removes the directory name from a file name
// It takes a string as an argument and returns a string
//
// name - the name of the file
func RemoveDirName(name string) string {
	return strings.Replace(name, "./"+constant.FILES_DIR, "", 1)
}
