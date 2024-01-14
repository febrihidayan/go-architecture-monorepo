package grpc_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (x *UserRepositoryMock) CreateUser(ctx context.Context, payload entities.User) (result *entities.User, err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.User); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *UserRepositoryMock) UpdateEmailVerifiedUser(ctx context.Context, payload entities.User) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
