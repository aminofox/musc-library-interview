package artist

import "github.com/gin-gonic/gin"

type DeleteArtistInput struct {
	ID string
}

type DeleteArtistOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func (t *artistUseCase) DeleteByID(ctx *gin.Context, input *DeleteArtistInput) (*DeleteArtistOutPut, error) {
	rowsAffected, err := t.artistRepository.Delete(ctx, input.ID)
	if err != nil {
		return &DeleteArtistOutPut{
			RowsAffected: 0,
		}, ErrDeleteArtistFailed
	}
	output := &DeleteArtistOutPut{
		RowsAffected: rowsAffected,
	}
	return output, nil
}
