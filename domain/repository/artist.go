package repository

import (
	"context"
	"music-libray-management/domain/entity"
)

type ArtistRepository interface {
	Create(ctx context.Context, input *entity.Artist) (string, error)
	Update(ctx context.Context, id string, input *entity.Artist) (int64, error)
	GetByID(ctx context.Context, id string) (*entity.Artist, error)
	GetList(ctx context.Context, params entity.GetListArtistOption) ([]*entity.Artist, error)
	Delete(ctx context.Context, id string) (int64, error)
}
