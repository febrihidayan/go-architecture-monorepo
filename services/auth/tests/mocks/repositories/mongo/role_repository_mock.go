package mongo_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

type RoleRepositoryMock struct {
	mock.Mock
}

func (x *RoleRepositoryMock) Create(ctx context.Context, payload *entities.Role) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *RoleRepositoryMock) Find(ctx context.Context, id string) (result *entities.Role, err error) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*entities.Role); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *RoleRepositoryMock) FindByName(ctx context.Context, name string) (result *entities.Role, err error) {
	args := x.Called(name)

	if n, ok := args.Get(0).(*entities.Role); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *RoleRepositoryMock) GetAll(ctx context.Context, params *entities.RoleQueryParams) (result []*entities.Role, total int, err error) {
	args := x.Called(params)

	if n, ok := args.Get(0).([]*entities.Role); ok {
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

func (x *RoleRepositoryMock) Update(ctx context.Context, payload *entities.Role) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
