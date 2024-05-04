package track

import "github.com/gin-gonic/gin"

type DeleteTrackInput struct {
	ID string
}

type DeleteTrackOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func (t *trackUseCase) DeleteByID(ctx *gin.Context, input *DeleteTrackInput) (*DeleteTrackOutPut, error) {
	rowsAffected, err := t.trackRepository.Delete(ctx, input.ID)
	if err != nil {
		return &DeleteTrackOutPut{
			RowsAffected: 0,
		}, ErrDeleteTrackFailed
	}
	output := &DeleteTrackOutPut{
		RowsAffected: rowsAffected,
	}
	return output, nil
}
