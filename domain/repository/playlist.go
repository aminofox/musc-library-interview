package repository

import (
	"context"
	"music-libray-management/domain/entity"
)

type PlaylistRepository interface {
	Create(ctx context.Context, input *entity.Playlist) (string, error)
	Update(ctx context.Context, id string, input *entity.Playlist) (int64, error)
	GetByID(ctx context.Context, id string) (*entity.Playlist, error)
	GetList(ctx context.Context, params entity.GetListPlaylistOption) ([]*entity.Playlist, error)
	Delete(ctx context.Context, id string) (int64, error)
}
