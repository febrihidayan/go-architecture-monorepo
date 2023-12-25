package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
)

type AwsService interface {
	UploadFile(ctx context.Context, payload *entities.File) (string, error)
	DeleteFile(url string) error
	ReadFile(fullPath string) (*s3.GetObjectOutput, error)
	GetObjects(ctx context.Context) error
}
