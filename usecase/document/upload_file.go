package document

import (
	"fmt"
	"mime"
	"mime/multipart"
	"music-libray-management/domain/entity"
	"music-libray-management/internal/helper"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type UploadFileInput struct {
	File     *multipart.FileHeader `form:"file" binding:"required"`
	Category string                `form:"category" binding:"required"`
	ParentID string                `form:"parent_id" binding:"required"`
	UserID   string
}

type UploadFileOutput struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

func (t *documentUseCase) UploadFile(ctx *gin.Context, input *UploadFileInput) (*UploadFileOutput, error) {
	content, fileName, _, err := helper.ReadFile(input.File)
	if err != nil {
		return nil, ErrReadFileFailed
	}
	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	path := fmt.Sprintf("%s/%s/%s", input.Category, input.UserID, fileName)

	publicURL, errUploadFile := t.storageRepository.UploadFile(strings.NewReader(string(content)), path, contentType)
	if errUploadFile != nil {
		return nil, errUploadFile
	}

	data := &entity.Document{
		Name:     fileName,
		Path:     publicURL,
		ParentID: input.ParentID, //albumID,userID,artistID
		Category: input.Category,
		Status:   entity.Completed,
	}

	id, err := t.documentRepository.Create(ctx, data)
	if err != nil {
		return nil, ErrUploadFailed
	}

	return &UploadFileOutput{
		id,
		publicURL,
	}, nil
}
