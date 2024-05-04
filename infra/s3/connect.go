package s3

import (
	"context"
	"music-libray-management/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	*s3.Client
}

func Connect(cfg *config.Environment) (*S3Client, error) {
	s3config, err := awsConfig.LoadDefaultConfig(
		context.Background(),
		awsConfig.WithRegion(cfg.AwsRegion),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AwsAccessKey, cfg.AwsSecret, ""),
		))
	if err != nil {
		return nil, err
	}
	s3Client := s3.NewFromConfig(s3config)
	params := &s3.ListObjectsInput{
		Bucket:  aws.String(cfg.S3Bucket),
		MaxKeys: aws.Int32(1),
	}
	_, err = s3Client.ListObjects(context.Background(), params)
	if err != nil {
		return nil, err
	}
	return &S3Client{
		Client: s3Client,
	}, nil
}
