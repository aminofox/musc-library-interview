package repository

import (
	"music-libray-management/config"
	"music-libray-management/domain/repository"
	"music-libray-management/infra/s3"
)

func NewStorageRepository(cfg *config.Environment, s3Client *s3.S3Client) repository.StorageRepository {
	return &storageRepository{
		cfg:      cfg,
		s3Client: s3Client,
	}
}

type storageRepository struct {
	cfg      *config.Environment
	s3Client *s3.S3Client
}
