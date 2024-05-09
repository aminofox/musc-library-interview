package repository

import (
	"context"
	"music-libray-management/domain/entity"
)

type AlbumRepository interface {
	Create(ctx context.Context, input *entity.Album) (string, error)
	Update(ctx context.Context, id string, input *entity.Album) (int64, error)
	GetByID(ctx context.Context, id string) (*entity.Album, error)
	GetByTitle(ctx context.Context, title string) ([]string, error)
	GetList(ctx context.Context, params entity.GetListAlbumOption) ([]*entity.Album, error)
	Delete(ctx context.Context, id string) (int64, error)
}
