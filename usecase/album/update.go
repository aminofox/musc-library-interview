package album

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UpdateAlbumInput struct {
	ID          string
	Title       string   `json:"title" bson:"title"`
	ReleaseYear int      `json:"release_year" bson:"release_year"`
	CoverImage  string   `json:"cover_image" bson:"cover_image"`
	Tracks      []string `json:"tracks" bson:"tracks"`
}

type UpdateAlbumOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func (t *albumUseCase) Update(ctx *gin.Context, input *UpdateAlbumInput) (*UpdateAlbumOutput, error) {
	data := &entity.Album{}

	err := copier.Copy(data, input)
	if err != nil {
		return &UpdateAlbumOutput{RowsAffected: 0}, ErrUpdateAlbumFailed
	}

	rowsAffected, err := t.albumRepository.Update(ctx, input.ID, data)
	if err != nil {
		return &UpdateAlbumOutput{RowsAffected: 0}, ErrUpdateAlbumFailed
	}

	return &UpdateAlbumOutput{
		rowsAffected,
	}, nil
}
