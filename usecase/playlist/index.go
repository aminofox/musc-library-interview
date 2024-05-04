package playlist

import (
	"errors"
	"music-libray-management/domain/entity"
	"music-libray-management/domain/repository"
	"music-libray-management/infra/mongo"

	"github.com/gin-gonic/gin"
)

var (
	ErrCreatePlaylistFailed   = errors.New("create playlist failed")
	ErrDeletePlaylistFailed   = errors.New("delete playlist failed")
	ErrGetByIdPlaylistFailed  = errors.New("get by id playlist failed")
	ErrGetListPlaylistFailed  = errors.New("get kist playlist failed")
	ErrUpdatePlaylistFailed   = errors.New("update playlist failed")
	ErrAddTrackPlaylistFailed = errors.New("update playlist failed")
)

type UseCase interface {
	Create(ctx *gin.Context, input *CreatePlaylistInput) (*CreatePlaylistOutput, error)
	Update(ctx *gin.Context, input *UpdatePlaylistInput) (*UpdatePlaylistOutput, error)
	DeleteByID(ctx *gin.Context, input *DeletePlaylistInput) (*DeletePlaylistOutPut, error)
	GetByID(ctx *gin.Context, input *GetPlaylistByIDInput) (*GetListPlaylistOutput, error)
	GetListPlaylist(ctx *gin.Context, input *GetListPlaylistInput) ([]*entity.Playlist, error)
	AddTrackToPlaylist(ctx *gin.Context, input *AddTrackToPlaylistInput) (*AddTrackToPlaylistOutput, error)
}

type playlistUseCase struct {
	playlistRepository repository.PlaylistRepository
	database           *mongo.MongoDB
}

func NewPlaylistUseCase(
	playlistRepository repository.PlaylistRepository,
	database *mongo.MongoDB,
) UseCase {
	return &playlistUseCase{
		playlistRepository: playlistRepository,
		database:           database,
	}
}
