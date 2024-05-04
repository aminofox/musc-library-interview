package album

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type CreateAlbumInput struct {
	Title       string   `json:"title" bson:"title"`
	ReleaseYear int      `json:"release_year" bson:"release_year"`
	CoverImage  string   `json:"cover_image" bson:"cover_image"`
	Tracks      []string `json:"tracks" bson:"tracks"`
}

type CreateAlbumOutput struct {
	ID string `json:"id"`
}

func (t *albumUseCase) Create(ctx *gin.Context, input *CreateAlbumInput) (*CreateAlbumOutput, error) {
	data := &entity.Album{}

	err := copier.Copy(data, input)
	if err != nil {
		return &CreateAlbumOutput{ID: ""}, ErrCreateAlbumFailed
	}

	id, err := t.albumRepository.Create(ctx, data)
	if err != nil {
		return &CreateAlbumOutput{ID: ""}, ErrCreateAlbumFailed
	}

	return &CreateAlbumOutput{
		ID: id,
	}, nil
}
