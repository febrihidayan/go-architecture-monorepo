package mongo_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/stretchr/testify/mock"
)

type CloudRepositoryMock struct {
	mock.Mock
}

func (x *CloudRepositoryMock) Create(ctx context.Context, payload *entities.Cloud) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *CloudRepositoryMock) All(ctx context.Context, params *entities.CloudQueryParams) (results []*entities.Cloud, err error) {
	args := x.Called(params)

	if n, ok := args.Get(0).([]*entities.Cloud); ok {
		results = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *CloudRepositoryMock) Find(ctx context.Context, id string) (result *entities.Cloud, err error) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*entities.Cloud); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *CloudRepositoryMock) FindByUrl(ctx context.Context, url string) (result *entities.Cloud, err error) {
	args := x.Called(url)

	if n, ok := args.Get(0).(*entities.Cloud); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *CloudRepositoryMock) Update(ctx context.Context, payload *entities.Cloud) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *CloudRepositoryMock) Delete(ctx context.Context, id string) (err error) {
	args := x.Called(id)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
