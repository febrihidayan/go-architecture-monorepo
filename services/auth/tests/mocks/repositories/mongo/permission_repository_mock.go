package mongo_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

type PermissionRepositoryMock struct {
	mock.Mock
}

func (x *PermissionRepositoryMock) Create(ctx context.Context, payload *entities.Permission) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *PermissionRepositoryMock) Find(ctx context.Context, id string) (result *entities.Permission, err error) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*entities.Permission); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *PermissionRepositoryMock) FindByName(ctx context.Context, name string) (result *entities.Permission, err error) {
	args := x.Called(name)

	if n, ok := args.Get(0).(*entities.Permission); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *PermissionRepositoryMock) GetAll(ctx context.Context, params *entities.PermissionQueryParams) (result []*entities.Permission, total int, err error) {
	args := x.Called(params)

	if n, ok := args.Get(0).([]*entities.Permission); ok {
		result = n
	}

	if n, ok := args.Get(1).(int); ok {
		total = n
	}

	if n, ok := args.Get(2).(error); ok {
		err = n
	}

	return
}

func (x *PermissionRepositoryMock) Update(ctx context.Context, payload *entities.Permission) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
