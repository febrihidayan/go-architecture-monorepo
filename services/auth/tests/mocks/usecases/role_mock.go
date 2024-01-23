package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

type RoleUsecaseMock struct {
	mock.Mock
}

func (x *RoleUsecaseMock) Create(ctx context.Context, payload entities.RoleDto) (result *entities.Role, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.Role); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *RoleUsecaseMock) Find(ctx context.Context, id string) (result *entities.Role, err *exceptions.CustomError) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*entities.Role); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *RoleUsecaseMock) GetAll(ctx context.Context, params entities.RoleQueryParams) (result *entities.RoleMeta, err *exceptions.CustomError) {
	args := x.Called(params)

	if n, ok := args.Get(0).(*entities.RoleMeta); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *RoleUsecaseMock) Update(ctx context.Context, payload entities.RoleDto) (result *entities.Role, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.Role); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}
