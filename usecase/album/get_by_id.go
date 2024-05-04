package album

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
)

type GetAlbumByIDInput struct {
	ID string
}

type GetListAlbumOutput struct {
	*entity.Album
}

func (t *albumUseCase) GetByID(ctx *gin.Context, input *GetAlbumByIDInput) (*GetListAlbumOutput, error) {
	album, err := t.albumRepository.GetByID(ctx, input.ID)
	if err != nil {
		return nil, ErrGetByIdAlbumFailed
	}
	return &GetListAlbumOutput{
		Album: album,
	}, nil
}
