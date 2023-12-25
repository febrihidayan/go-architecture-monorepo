package services

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
)

type AwsService struct {
	cfg *config.AwsConfig
	ctx time.Duration
}

type BucketResponse struct {
	PathUrl string
}

func (s *AwsService) Client() (*s3.Client, error) {
	client := s3.New(s3.Options{
		Region: string(s.cfg.Region),
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			s.cfg.AccessKey,
			s.cfg.AccessSecret,
			"",
		)),
	})

	return client, nil
}

func (s *AwsService) GetPathDefault(fullPath string) string {
	return fmt.Sprintf(
		"https://%s.s3.%s.amazonaws.com/%s",
		s.cfg.Bucket,
		string(s.cfg.Region),
		fullPath,
	)
}

func (s *AwsService) UploadFile(ctx context.Context, payload *entities.File) (string, error) {
	bucket, err := s.Client()
	if err != nil {
		return "", err
	}

	file, err := os.Open(payload.Name)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fullPath := fmt.Sprintf("%s/%s", payload.Directory, strings.ReplaceAll(payload.Origin, " ", "-"))

	_, errUpload := bucket.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(s.cfg.Bucket),
		Key:           aws.String(fullPath),
		Body:          file,
		ContentLength: aws.Int64(int64(payload.Size)),
		ACL:           types.ObjectCannedACLPublicRead, // by default if acl is not set, the object is private
		ContentType:   aws.String(payload.ContentType),
	})

	if errUpload != nil {
		return "", errUpload
	}

	return s.GetPathDefault(fullPath), nil
}

func (s *AwsService) DeleteFile(fullPath string) error {
	bucket, err := s.Client()
	if err != nil {
		return err
	}

	path := strings.Replace(fullPath, s.GetPathDefault(fullPath), "", 1)

	_, errDelete := bucket.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.cfg.Bucket),
		Key:    aws.String(path),
	})

	if errDelete != nil {
		return errDelete
	}

	return nil
}

func (s *AwsService) ReadFile(fullPath string) (*s3.GetObjectOutput, error) {
	bucket, err := s.Client()
	if err != nil {
		return nil, err
	}

	path := strings.Replace(fullPath, s.GetPathDefault(fullPath), "", 1)

	result, errObject := bucket.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.cfg.Bucket),
		Key:    aws.String(path),
	})

	if errObject != nil {
		return nil, errObject
	}

	return result, nil
}

func (s *AwsService) GetObjects(ctx context.Context) error {
	bucket, err := s.Client()
	if err != nil {
		return err
	}
	lsRes, errLs := bucket.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(s.cfg.Bucket),
	})
	if errLs != nil {
		return errLs
	}

	for _, object := range lsRes.Contents {
		fmt.Println("Objects:", object.Key)
	}
	return nil
}

func NewAwsService(cfg *config.AwsConfig) *AwsService {
	return &AwsService{
		cfg: cfg,
		ctx: cfg.RequestTimeout,
	}
}
