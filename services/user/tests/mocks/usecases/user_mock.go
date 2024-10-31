package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (x *UserUsecaseMock) Create(ctx context.Context, payload entities.UserDto) (result *entities.User, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.User); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *UserUsecaseMock) CreateAuth(ctx context.Context, payload entities.UserDto) (result *entities.User, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.User); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *UserUsecaseMock) Find(ctx context.Context, id string) (result *entities.User, err *exceptions.CustomError) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*entities.User); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *UserUsecaseMock) GetAll(ctx context.Context, params entities.UserQueryParams) (result *entities.UserMeta, err *exceptions.CustomError) {
	args := x.Called(params)

	if n, ok := args.Get(0).(*entities.UserMeta); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *UserUsecaseMock) Update(ctx context.Context, payload entities.UserDto) (result *entities.User, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.User); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *UserUsecaseMock) UpdateEmailVerified(ctx context.Context, payload entities.UserDto) (err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *UserUsecaseMock) Delete(ctx context.Context, id string) (err *exceptions.CustomError) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}
