package album

import "github.com/gin-gonic/gin"

type DeleteAlbumInput struct {
	ID string
}

type DeleteAlbumOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func (t *albumUseCase) DeleteByID(ctx *gin.Context, input *DeleteAlbumInput) (*DeleteAlbumOutPut, error) {
	rowsAffected, err := t.albumRepository.Delete(ctx, input.ID)
	if err != nil {
		return &DeleteAlbumOutPut{
			RowsAffected: 0,
		}, ErrDeleteAlbumFailed
	}
	output := &DeleteAlbumOutPut{
		RowsAffected: rowsAffected,
	}
	return output, nil
}
