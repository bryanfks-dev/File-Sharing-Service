package utils

import (
	"main/core/enum"
	"slices"
)

// GetFileType is a function that returns the file type based on the extension
// of the file.
//
// extension - the extension of the file
func GetFileType(extension string) enum.Type {
	if slices.Contains(enum.Music.Extensions(), extension) {
		return enum.Music
	}

	if slices.Contains(enum.Image.Extensions(), extension) {
		return enum.Image
	}

	if slices.Contains(enum.Video.Extensions(), extension) {
		return enum.Video
	}

	if slices.Contains(enum.Document.Extensions(), extension) {
		return enum.Document
	}

	if slices.Contains(enum.Archive.Extensions(), extension) {
		return enum.Archive
	}

	return enum.Other
}
