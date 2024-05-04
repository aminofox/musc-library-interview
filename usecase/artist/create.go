package artist

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type CreateArtistInput struct {
	Name    string   `json:"name" bson:"name"`
	Country string   `json:"country" bson:"country"`
	Avatar  string   `json:"avatar" bson:"avatar"`
	Tracks  []string `json:"tracks" bson:"tracks"`
}

type CreateArtistOutput struct {
	ID string `json:"id"`
}

func (t *artistUseCase) Create(ctx *gin.Context, input *CreateArtistInput) (*CreateArtistOutput, error) {
	data := &entity.Artist{}

	err := copier.Copy(data, input)
	if err != nil {
		return &CreateArtistOutput{ID: ""}, ErrCreateArtistFailed
	}

	id, err := t.artistRepository.Create(ctx, data)
	if err != nil {
		return &CreateArtistOutput{ID: ""}, ErrCreateArtistFailed
	}

	return &CreateArtistOutput{
		ID: id,
	}, nil
}
