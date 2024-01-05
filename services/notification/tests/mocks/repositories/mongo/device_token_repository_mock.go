package mongo_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/stretchr/testify/mock"
)

type DeviceTokenRepositoryMock struct {
	mock.Mock
}

func (x *DeviceTokenRepositoryMock) Create(ctx context.Context, payload *entities.DeviceToken) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *DeviceTokenRepositoryMock) All(ctx context.Context, params *entities.DeviceTokenQueryParams) (result []*entities.DeviceToken, err error) {
	args := x.Called(params)

	if n, ok := args.Get(0).([]*entities.DeviceToken); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *DeviceTokenRepositoryMock) Delete(ctx context.Context, id string) (err error) {
	args := x.Called(id)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
