package repository

import "strings"

type StorageRepository interface {
	UploadFile(file *strings.Reader, filename, contentType string) (path string, err error)
}
