package track

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
)

type GetTrackByIDInput struct {
	ID string
}

type GetListTrackOutput struct {
	*entity.Track
}

func (t *trackUseCase) GetByID(ctx *gin.Context, input *GetTrackByIDInput) (*GetListTrackOutput, error) {
	track, err := t.trackRepository.GetByID(ctx, input.ID)
	if err != nil {
		return nil, ErrGetByIdTrackFailed
	}
	return &GetListTrackOutput{
		Track: track,
	}, nil
}
