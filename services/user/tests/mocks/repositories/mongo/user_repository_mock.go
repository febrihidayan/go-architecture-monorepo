package mongo_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (x *UserRepositoryMock) Create(ctx context.Context, payload *entities.User) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *UserRepositoryMock) Find(ctx context.Context, id string) (result *entities.User, err error) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*entities.User); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *UserRepositoryMock) FindByEmail(ctx context.Context, email string) (result *entities.User, err error) {
	args := x.Called(email)

	if n, ok := args.Get(0).(*entities.User); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *UserRepositoryMock) GetAll(ctx context.Context, params *entities.UserQueryParams) (result []*entities.User, total int, err error) {
	args := x.Called(params)

	if n, ok := args.Get(0).([]*entities.User); ok {
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

func (x *UserRepositoryMock) Update(ctx context.Context, payload *entities.User) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
