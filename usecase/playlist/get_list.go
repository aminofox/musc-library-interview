package playlist

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
)

type GetListPlaylistInput struct {
	PageIndex int    `form:"pageIndex"`
	PageSize  int    `form:"pageSize"`
	Order     string `form:"order"`
	Name      string `form:"name"`
	Track     string `form:"track"`
}

func (t *playlistUseCase) GetListPlaylist(ctx *gin.Context, input *GetListPlaylistInput) ([]*entity.Playlist, error) {
	tracks, err := t.playlistRepository.GetList(ctx, entity.GetListPlaylistOption{
		GetListOption: entity.GetListOption{
			PageIndex: input.PageIndex,
			PageSize:  input.PageSize,
			Order:     input.Order,
		},
		Name:  input.Name,
		Track: input.Track,
	})
	if err != nil {
		return nil, ErrGetListPlaylistFailed
	}
	return tracks, nil
}
