package fileutil

// ImageMimeTypes refer net/http package
var ImageMimeTypes = map[string]string{
	"bmp": "image/bmp",
	"gif": "image/gif",
	"ief": "image/ief",
	"jpg": "image/jpeg",
	// "jpe":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"svg":  "image/svg+xml",
	"ico":  "image/x-icon",
	"webp": "image/webp",
}

// IsImageFile check file is image file.
func IsImageFile(path string) bool {
	mime := MimeType(path)
	if mime == "" {
		return false
	}

	for _, imgMime := range ImageMimeTypes {
		if imgMime == mime {
			return true
		}
	}
	return false
}