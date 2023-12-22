package mongo_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

type AuthRepositoryMock struct {
	mock.Mock
}

func (x *AuthRepositoryMock) Create(ctx context.Context, payload *entities.Auth) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *AuthRepositoryMock) Find(ctx context.Context, id string) (result *entities.Auth, err error) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*entities.Auth); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *AuthRepositoryMock) FindByEmail(ctx context.Context, email string) (result *entities.Auth, err error) {
	args := x.Called(email)

	if n, ok := args.Get(0).(*entities.Auth); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *AuthRepositoryMock) FindByUserId(ctx context.Context, userId string) (result *entities.Auth, err error) {
	args := x.Called(userId)

	if n, ok := args.Get(0).(*entities.Auth); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *AuthRepositoryMock) Update(ctx context.Context, payload *entities.Auth) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
