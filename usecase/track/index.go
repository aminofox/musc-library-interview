package track

import (
	"errors"
	"music-libray-management/domain/entity"
	"music-libray-management/domain/repository"
	"music-libray-management/infra/mongo"

	"github.com/gin-gonic/gin"
)

var (
	ErrCreateTrackFailed  = errors.New("create track failed")
	ErrDeleteTrackFailed  = errors.New("delete track failed")
	ErrGetByIdTrackFailed = errors.New("get by id track failed")
	ErrGetListTrackFailed = errors.New("get list track failed")
	ErrUpdateTrackFailed  = errors.New("update track failed")
)

type UseCase interface {
	Create(ctx *gin.Context, input *CreateTrackInput) (*CreateTrackOutput, error)
	Update(ctx *gin.Context, input *UpdateTrackInput) (*UpdateTrackOutput, error)
	DeleteByID(ctx *gin.Context, input *DeleteTrackInput) (*DeleteTrackOutPut, error)
	GetByID(ctx *gin.Context, input *GetTrackByIDInput) (*GetListTrackOutput, error)
	GetListTrack(ctx *gin.Context, input *GetListTrackInput) ([]*entity.Track, error)
}

type trackUseCase struct {
	trackRepository  repository.TrackRepository
	artistRepository repository.ArtistRepository
	albumRepository  repository.AlbumRepository
	database         *mongo.MongoDB
}

func NewTrackUseCase(
	trackRepository repository.TrackRepository,
	artistRepository repository.ArtistRepository,
	albumRepository repository.AlbumRepository,
	database *mongo.MongoDB,
) UseCase {
	return &trackUseCase{
		trackRepository:  trackRepository,
		artistRepository: artistRepository,
		albumRepository:  albumRepository,
		database:         database,
	}
}
