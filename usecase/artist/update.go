package artist

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UpdateArtistInput struct {
	ID      string
	Name    string   `json:"name" bson:"name"`
	Country string   `json:"country" bson:"country"`
	Avatar  string   `json:"avatar" bson:"avatar"`
	Tracks  []string `json:"tracks" bson:"tracks"`
}

type UpdateArtistOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func (t *artistUseCase) Update(ctx *gin.Context, input *UpdateArtistInput) (*UpdateArtistOutput, error) {
	data := &entity.Artist{}

	err := copier.Copy(data, input)
	if err != nil {
		return &UpdateArtistOutput{RowsAffected: 0}, ErrUpdateArtistFailed
	}

	rowsAffected, err := t.artistRepository.Update(ctx, input.ID, data)
	if err != nil {
		return &UpdateArtistOutput{RowsAffected: 0}, ErrUpdateArtistFailed
	}

	return &UpdateArtistOutput{
		rowsAffected,
	}, nil
}
