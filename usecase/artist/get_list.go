package artist

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
)

type GetListArtistInput struct {
	PageIndex int    `form:"pageIndex"`
	PageSize  int    `form:"pageSize"`
	Order     string `form:"order"`
	Name      string `form:"name"`
	Track     string `form:"track"`
}

func (t *artistUseCase) GetListArtist(ctx *gin.Context, input *GetListArtistInput) ([]*entity.Artist, error) {
	artists, err := t.artistRepository.GetList(ctx, entity.GetListArtistOption{
		GetListOption: entity.GetListOption{
			PageIndex: input.PageIndex,
			PageSize:  input.PageSize,
			Order:     input.Order,
		},
		Name:  input.Name,
		Track: input.Track,
	})
	if err != nil {
		return nil, ErrGetListArtistFailed
	}
	return artists, nil
}
