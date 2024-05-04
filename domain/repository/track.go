package repository

import (
	"context"
	"music-libray-management/domain/entity"
)

type TrackRepository interface {
	Create(ctx context.Context, input *entity.Track) (string, error)
	Update(ctx context.Context, id string, input *entity.Track) (int64, error)
	GetByID(ctx context.Context, id string) (*entity.Track, error)
	GetList(ctx context.Context, params entity.GetListTrackOption) ([]*entity.Track, error)
	Delete(ctx context.Context, id string) (int64, error)
}
