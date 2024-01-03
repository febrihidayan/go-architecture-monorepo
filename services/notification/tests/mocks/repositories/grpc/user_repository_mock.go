package grpc_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (x *UserRepositoryMock) FindUser(ctx context.Context, id string) (result *entities.User, err error) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*entities.User); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}
