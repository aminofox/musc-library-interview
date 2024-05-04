package document

import (
	"errors"
	"music-libray-management/domain/repository"
	"music-libray-management/infra/mongo"

	"github.com/gin-gonic/gin"
)

var (
	ErrUploadFailed   = errors.New("upload file failed")
	ErrReadFileFailed = errors.New("can not read file")
)

type UseCase interface {
	UploadFile(ctx *gin.Context, input *UploadFileInput) (*UploadFileOutput, error)
}

type documentUseCase struct {
	documentRepository repository.DocumentRepository
	storageRepository  repository.StorageRepository
	database           *mongo.MongoDB
}

func NewDocumentUseCase(
	documentRepository repository.DocumentRepository,
	storageRepository repository.StorageRepository,
	database *mongo.MongoDB,
) UseCase {
	return &documentUseCase{
		documentRepository: documentRepository,
		storageRepository:  storageRepository,
		database:           database,
	}
}
