package grpc_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/stretchr/testify/mock"
)

type AuthRepositoryMock struct {
	mock.Mock
}

func (x *AuthRepositoryMock) CreateOrUpdate(ctx context.Context, payload *entities.Auth) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
