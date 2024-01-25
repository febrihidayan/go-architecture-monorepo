package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/stretchr/testify/mock"
)

type ProfileUsecaseMock struct {
	mock.Mock
}

func (x *ProfileUsecaseMock) Find(ctx context.Context, id string) (result *entities.User, err *exceptions.CustomError) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*entities.User); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *ProfileUsecaseMock) Update(ctx context.Context, payload entities.UserDto) (result *entities.User, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.User); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}
