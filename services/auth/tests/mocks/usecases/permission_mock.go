package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

type PermissionUsecaseMock struct {
	mock.Mock
}

func (x *PermissionUsecaseMock) Create(ctx context.Context, payload entities.PermissionDto) (result *entities.Permission, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.Permission); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *PermissionUsecaseMock) Find(ctx context.Context, id string) (result *entities.Permission, err *exceptions.CustomError) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*entities.Permission); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *PermissionUsecaseMock) GetAll(ctx context.Context, params entities.PermissionQueryParams) (result *entities.PermissionMeta, err *exceptions.CustomError) {
	args := x.Called(params)

	if n, ok := args.Get(0).(*entities.PermissionMeta); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *PermissionUsecaseMock) Update(ctx context.Context, payload entities.PermissionDto) (result *entities.Permission, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.Permission); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}
