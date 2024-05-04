package track

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
)

type GetListTrackInput struct {
	PageIndex int    `form:"pageIndex"`
	PageSize  int    `form:"pageSize"`
	Order     string `form:"order"`
	Title     string `form:"title"`
	Artist    string `form:"artist"`
	Album     string `form:"album"`
	Genre     string `form:"genre"`
}

func (t *trackUseCase) GetListTrack(ctx *gin.Context, input *GetListTrackInput) ([]*entity.Track, error) {
	tracks, err := t.trackRepository.GetList(ctx, entity.GetListTrackOption{
		GetListOption: entity.GetListOption{
			PageIndex: input.PageIndex,
			PageSize:  input.PageSize,
			Order:     input.Order,
		},
		Title:  input.Title,
		Artist: input.Artist,
		Album:  input.Album,
		Genre:  input.Genre,
	})
	if err != nil {
		return nil, ErrGetListTrackFailed
	}
	return tracks, nil
}
