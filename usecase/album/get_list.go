package album

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
)

type GetListAlbumInput struct {
	PageIndex   int    `form:"pageIndex"`
	PageSize    int    `form:"pageSize"`
	Order       string `form:"order"`
	Title       string `form:"title"`
	ReleaseYear int    `form:"release_year"`
	Track       string `form:"track"`
}

func (t *albumUseCase) GetListAlbum(ctx *gin.Context, input *GetListAlbumInput) ([]*entity.Album, error) {
	albums, err := t.albumRepository.GetList(ctx, entity.GetListAlbumOption{
		GetListOption: entity.GetListOption{
			PageIndex: input.PageIndex,
			PageSize:  input.PageSize,
			Order:     input.Order,
		},
		Title:       input.Title,
		ReleaseYear: input.ReleaseYear,
		Track:       input.Track,
	})
	if err != nil {
		return nil, ErrGetListAlbumFailed
	}
	return albums, nil
}
