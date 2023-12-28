package grpc_repositories

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type StorageRepositoryMock struct {
	mock.Mock
}

func (x *StorageRepositoryMock) UpdateCloudApprove(ctx context.Context, url []string) (err error) {
	args := x.Called(url)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *StorageRepositoryMock) DeleteCloudApprove(ctx context.Context, url []string) (err error) {
	args := x.Called(url)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
