package repository

import (
	"context"
	"music-libray-management/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, input *entity.User) (string, error)
	Update(ctx context.Context, id string, input *entity.User) (string, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetList(ctx context.Context) ([]*entity.User, error)
	Delete(ctx context.Context, id string) (int64, error)
}
