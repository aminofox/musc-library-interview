package playlist

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type CreatePlaylistInput struct {
	Name   string   `json:"name" bson:"name"`
	Tracks []string `json:"tracks" bson:"tracks"`
}

type CreatePlaylistOutput struct {
	ID string `json:"id"`
}

func (t *playlistUseCase) Create(ctx *gin.Context, input *CreatePlaylistInput) (*CreatePlaylistOutput, error) {
	data := &entity.Playlist{}

	err := copier.Copy(data, input)
	if err != nil {
		return &CreatePlaylistOutput{ID: ""}, ErrCreatePlaylistFailed
	}

	id, err := t.playlistRepository.Create(ctx, data)
	if err != nil {
		return &CreatePlaylistOutput{ID: ""}, ErrCreatePlaylistFailed
	}

	return &CreatePlaylistOutput{
		ID: id,
	}, nil
}
