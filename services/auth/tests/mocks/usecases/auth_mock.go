package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

type AuthUsecaseMock struct {
	mock.Mock
}

func (x *AuthUsecaseMock) Login(ctx context.Context, payload entities.AuthDto) (result *entities.AuthTokenMeta, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.AuthTokenMeta); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AuthUsecaseMock) Register(ctx context.Context, payload entities.RegisterDto) (result *entities.Auth, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.Auth); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AuthUsecaseMock) CreateOrUpdate(ctx context.Context, payload entities.AuthDto) (result *entities.Auth, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.Auth); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AuthUsecaseMock) EmailVerified(ctx context.Context, token string) (err *exceptions.CustomError) {
	args := x.Called(token)

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AuthUsecaseMock) SendEmailVerified(ctx context.Context, email string) (err *exceptions.CustomError) {
	args := x.Called(email)

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AuthUsecaseMock) PasswordEmail(ctx context.Context, email string) (err *exceptions.CustomError) {
	args := x.Called(email)

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *AuthUsecaseMock) PasswordReset(ctx context.Context, payload entities.PasswordReset) (err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}
