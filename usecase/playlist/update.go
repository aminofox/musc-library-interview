package playlist

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UpdatePlaylistInput struct {
	ID     string
	Name   string   `json:"name" bson:"name"`
	Tracks []string `json:"tracks" bson:"tracks"`
}

type UpdatePlaylistOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func (t *playlistUseCase) Update(ctx *gin.Context, input *UpdatePlaylistInput) (*UpdatePlaylistOutput, error) {
	data := &entity.Playlist{}

	err := copier.Copy(data, input)
	if err != nil {
		return &UpdatePlaylistOutput{RowsAffected: 0}, ErrUpdatePlaylistFailed
	}

	rowsAffected, err := t.playlistRepository.Update(ctx, input.ID, data)
	if err != nil {
		return &UpdatePlaylistOutput{RowsAffected: 0}, ErrUpdatePlaylistFailed
	}

	return &UpdatePlaylistOutput{
		rowsAffected,
	}, nil
}
