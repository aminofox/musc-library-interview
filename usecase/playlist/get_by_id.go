package playlist

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
)

type GetPlaylistByIDInput struct {
	ID string
}

type GetListPlaylistOutput struct {
	*entity.Playlist
}

func (t *playlistUseCase) GetByID(ctx *gin.Context, input *GetPlaylistByIDInput) (*GetListPlaylistOutput, error) {
	track, err := t.playlistRepository.GetByID(ctx, input.ID)
	if err != nil {
		return nil, ErrGetByIdPlaylistFailed
	}
	return &GetListPlaylistOutput{
		Playlist: track,
	}, nil
}
