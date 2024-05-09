package track

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
)

type GetListTrackInput struct {
	PageIndex  int     `form:"pageIndex"`
	PageSize   int     `form:"pageSize"`
	Order      string  `form:"order"`
	Title      string  `form:"title"`
	Artist     string  `form:"artist"`
	Album      string  `form:"album"`
	Genre      string  `form:"genre"`
	ArtistName *string `form:"artist_name"`
	AlbumName  *string `form:"album_name"`
}

func (t *trackUseCase) GetListTrack(ctx *gin.Context, input *GetListTrackInput) ([]*entity.Track, error) {
	var artistIds []string
	var albumIds []string
	var err error
	if input.ArtistName != nil {
		artistIds, err = t.artistRepository.GetByName(ctx, *input.ArtistName)
		if err != nil {
			return nil, ErrGetListTrackFailed
		}
	}

	if input.AlbumName != nil {
		albumIds, err = t.albumRepository.GetByTitle(ctx, *input.AlbumName)
		if err != nil {
			return nil, ErrGetListTrackFailed
		}
	}

	tracks, err := t.trackRepository.GetList(ctx, entity.GetListTrackOption{
		GetListOption: entity.GetListOption{
			PageIndex: input.PageIndex,
			PageSize:  input.PageSize,
			Order:     input.Order,
		},
		Title:     input.Title,
		Artist:    input.Artist,
		Album:     input.Album,
		Genre:     input.Genre,
		ArtistIds: &artistIds,
		AlbumIds:  &albumIds,
	})
	if err != nil {
		return nil, ErrGetListTrackFailed
	}
	return tracks, nil
}
