package artist

import (
	"errors"
	"music-libray-management/domain/entity"
	"music-libray-management/domain/repository"
	"music-libray-management/infra/mongo"

	"github.com/gin-gonic/gin"
)

var (
	ErrCreateArtistFailed  = errors.New("create artist failed")
	ErrDeleteArtistFailed  = errors.New("delete artist failed")
	ErrGetByIdArtistFailed = errors.New("get by id artist failed")
	ErrGetListArtistFailed = errors.New("get kist artist failed")
	ErrUpdateArtistFailed  = errors.New("update artist failed")
)

type UseCase interface {
	Create(ctx *gin.Context, input *CreateArtistInput) (*CreateArtistOutput, error)
	Update(ctx *gin.Context, input *UpdateArtistInput) (*UpdateArtistOutput, error)
	DeleteByID(ctx *gin.Context, input *DeleteArtistInput) (*DeleteArtistOutPut, error)
	GetByID(ctx *gin.Context, input *GetArtistByIDInput) (*GetListArtistOutput, error)
	GetListArtist(ctx *gin.Context, input *GetListArtistInput) ([]*entity.Artist, error)
}

type artistUseCase struct {
	artistRepository repository.ArtistRepository
	database         *mongo.MongoDB
}

func NewArtistUseCase(
	artistRepository repository.ArtistRepository,
	database *mongo.MongoDB,
) UseCase {
	return &artistUseCase{
		artistRepository: artistRepository,
		database:         database,
	}
}
