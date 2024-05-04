package album

import (
	"errors"
	"music-libray-management/domain/entity"
	"music-libray-management/domain/repository"
	"music-libray-management/infra/mongo"

	"github.com/gin-gonic/gin"
)

var (
	ErrCreateAlbumFailed  = errors.New("create album failed")
	ErrDeleteAlbumFailed  = errors.New("delete album failed")
	ErrGetByIdAlbumFailed = errors.New("get by id album failed")
	ErrGetListAlbumFailed = errors.New("get kist album failed")
	ErrUpdateAlbumFailed  = errors.New("update album failed")
)

type UseCase interface {
	Create(ctx *gin.Context, input *CreateAlbumInput) (*CreateAlbumOutput, error)
	Update(ctx *gin.Context, input *UpdateAlbumInput) (*UpdateAlbumOutput, error)
	DeleteByID(ctx *gin.Context, input *DeleteAlbumInput) (*DeleteAlbumOutPut, error)
	GetByID(ctx *gin.Context, input *GetAlbumByIDInput) (*GetListAlbumOutput, error)
	GetListAlbum(ctx *gin.Context, input *GetListAlbumInput) ([]*entity.Album, error)
}

type albumUseCase struct {
	albumRepository repository.AlbumRepository
	database        *mongo.MongoDB
}

func NewAlbumUseCase(
	albumRepository repository.AlbumRepository,
	database *mongo.MongoDB,
) UseCase {
	return &albumUseCase{
		albumRepository: albumRepository,
		database:        database,
	}
}
