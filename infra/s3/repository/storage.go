package repository

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
)

func (sr *storageRepository) UploadFile(file *strings.Reader, filename, contentType string) (string, error) {
	if strings.Contains(filename, ".pdf") {
		contentType = "application/pdf"
	}

	objectInput := &s3.PutObjectInput{
		Bucket:      aws.String(sr.cfg.S3Bucket),
		Key:         aws.String(filename),
		Body:        file,
		ContentType: &contentType,
		ACL:         types.ObjectCannedACLPublicRead,
	}
	uploader := manager.NewUploader(sr.s3Client)
	result, err := uploader.Upload(context.Background(), objectInput)
	if err != nil {
		return "", err
	}
	return result.Location, nil
}
