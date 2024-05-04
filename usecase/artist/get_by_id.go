package artist

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
)

type GetArtistByIDInput struct {
	ID string
}

type GetListArtistOutput struct {
	*entity.Artist
}

func (t *artistUseCase) GetByID(ctx *gin.Context, input *GetArtistByIDInput) (*GetListArtistOutput, error) {
	artist, err := t.artistRepository.GetByID(ctx, input.ID)
	if err != nil {
		return nil, ErrGetByIdArtistFailed
	}
	return &GetListArtistOutput{
		Artist: artist,
	}, nil
}
