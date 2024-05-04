package repository

import (
	"context"
	"music-libray-management/domain/entity"
)

type DocumentRepository interface {
	Create(ctx context.Context, input *entity.Document) (string, error)
	Update(ctx context.Context, id string, input *entity.Document) (int64, error)
	GetByID(ctx context.Context, id string) (*entity.Document, error)
	GetList(ctx context.Context) ([]*entity.Document, error)
	Delete(ctx context.Context, id string) (int64, error)
}
