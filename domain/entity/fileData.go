package entity

// FileData is a struct that represents the file data
type FileData struct {
	FileName   string `json:"fileName"`
	FileSize   int64  `json:"fileSize"`
	FileType   string `json:"fileType"`
	UploadTime string `json:"uploadTime"`
	FileURL    string `json:"fileURL"`
}
