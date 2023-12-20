package mongo_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

type RoleUserRepositoryMock struct {
	mock.Mock
}

func (x *RoleUserRepositoryMock) CreateMany(ctx context.Context, payload []*entities.RoleUser) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *RoleUserRepositoryMock) AllByUserId(ctx context.Context, userId string) (result []*entities.RoleUser, err error) {
	args := x.Called(userId)

	if n, ok := args.Get(0).([]*entities.RoleUser); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *RoleUserRepositoryMock) Delete(ctx context.Context, payload *entities.RoleUser) (err error) {
	args := x.Called(payload.UserId, payload.RoleId)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}