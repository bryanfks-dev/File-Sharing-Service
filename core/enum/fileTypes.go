package enum

type Type int

const (
	Music Type = iota
	Image
	Video
	Document
	Archive
	Other
)

// Label is a method that returns the label of the enum
func (t Type) Label() string {
	switch t {
	case Music:
		return "Music"
	case Image:
		return "Image"
	case Video:
		return "Video"
	case Document:
		return "Document"
	case Archive:
		return "Archive"
	case Other:
		return "Other"
	default:
		return "Unknown"
	}
}

// Extensions is a method that returns the extensions of the enum
func (t Type) Extensions() []string {
	switch t {
	case Music:
		return []string{"mp3", "flac", "wav", "ogg", "m4a"}
	case Image:
		return []string{"jpg", "jpeg", "png", "gif", "bmp", "svg", "webp"}
	case Video:
		return []string{"mp4", "mkv", "avi", "mov", "flv", "wmv", "webm"}
	case Document:
		return []string{"pdf", "doc", "docx", "xls", "xlsx", "ppt", "pptx", "txt", "rtf"}
	case Archive:
		return []string{"zip", "rar", "tar", "gz", "7z"}
	default:
		return []string{}
	}
}
