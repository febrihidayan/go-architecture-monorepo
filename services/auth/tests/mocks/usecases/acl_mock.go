package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

type AclUsecaseMock struct {
	mock.Mock
}

func (x *AclUsecaseMock) GetAllPermissionByRole(ctx context.Context, roleId string) (result []*entities.Permission, err *exceptions.CustomError) {
	args := x.Called(roleId)

	if n, ok := args.Get(0).([]*entities.Permission); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AclUsecaseMock) GetAllPermission(ctx context.Context) (result []*entities.Permission, err *exceptions.CustomError) {
	args := x.Called()

	if n, ok := args.Get(0).([]*entities.Permission); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AclUsecaseMock) GetAllRole(ctx context.Context) (result []*entities.Role, err *exceptions.CustomError) {
	args := x.Called()

	if n, ok := args.Get(0).([]*entities.Role); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AclUsecaseMock) AccessUserLogin(ctx context.Context, userId string) (result *entities.AclMeta, err *exceptions.CustomError) {
	args := x.Called(userId)

	if n, ok := args.Get(0).(*entities.AclMeta); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AclUsecaseMock) GetAllUser(ctx context.Context, userId string) (result *entities.AclMeta, err *exceptions.CustomError) {
	args := x.Called(userId)

	if n, ok := args.Get(0).(*entities.AclMeta); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AclUsecaseMock) UpdatePermissionByRole(ctx context.Context, payload entities.AclPermissionDto) (err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AclUsecaseMock) UpdateUser(ctx context.Context, payload entities.AclUserDto) (err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}
