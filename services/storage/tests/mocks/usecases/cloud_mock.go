package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/stretchr/testify/mock"
)

type CloudUsecaseMock struct {
	mock.Mock
}

func (x *CloudUsecaseMock) Create(ctx context.Context, payload entities.CloudDto) (result *entities.Cloud, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.Cloud); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *CloudUsecaseMock) DeleteAllStatusJob(ctx context.Context) (err *exceptions.CustomError) {
	args := x.Called()

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *CloudUsecaseMock) Deletes(ctx context.Context, payloads []*entities.Cloud) (err *exceptions.CustomError) {
	args := x.Called(payloads)

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *CloudUsecaseMock) UpdateStatus(ctx context.Context, payloads []*entities.Cloud) (err *exceptions.CustomError) {
	args := x.Called(payloads)

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}
