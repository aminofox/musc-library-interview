package playlist

import "github.com/gin-gonic/gin"

type DeletePlaylistInput struct {
	ID string
}

type DeletePlaylistOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func (t *playlistUseCase) DeleteByID(ctx *gin.Context, input *DeletePlaylistInput) (*DeletePlaylistOutPut, error) {
	rowsAffected, err := t.playlistRepository.Delete(ctx, input.ID)
	if err != nil {
		return &DeletePlaylistOutPut{
			RowsAffected: 0,
		}, ErrDeletePlaylistFailed
	}
	output := &DeletePlaylistOutPut{
		RowsAffected: rowsAffected,
	}
	return output, nil
}
