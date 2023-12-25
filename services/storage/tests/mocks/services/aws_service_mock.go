package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/stretchr/testify/mock"
)

type AwsServiceMock struct {
	mock.Mock
}

func (m *AwsServiceMock) UploadFile(ctx context.Context, payload *entities.File) (string, error) {
	arg := m.Called(payload)
	var (
		link string
		err  error
	)

	if n, ok := arg.Get(0).(string); ok {
		link = n
	}
	if n, ok := arg.Get(1).(error); ok {
		err = n
	}

	return link, err
}

func (m *AwsServiceMock) DeleteFile(url string) error {
	arg := m.Called(url)
	var err error

	if n, ok := arg.Get(0).(error); ok {
		err = n
	}

	return err
}

func (m *AwsServiceMock) ReadFile(fullPath string) (*s3.GetObjectOutput, error) {
	arg := m.Called(fullPath)
	var (
		object *s3.GetObjectOutput
		err    error
	)

	if n, ok := arg.Get(0).(*s3.GetObjectOutput); ok {
		object = n
	}
	if n, ok := arg.Get(1).(error); ok {
		err = n
	}

	return object, err
}

func (m *AwsServiceMock) GetObjects(ctx context.Context) error {
	arg := m.Called()
	var err error

	if n, ok := arg.Get(0).(error); ok {
		err = n
	}

	return err
}
