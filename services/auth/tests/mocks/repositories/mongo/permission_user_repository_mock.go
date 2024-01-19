package mongo_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

type PermissionUserRepositoryMock struct {
	mock.Mock
}

func (x *PermissionUserRepositoryMock) CreateMany(ctx context.Context, payload []*entities.PermissionUser) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *PermissionUserRepositoryMock) AllByUserId(ctx context.Context, userId string) (result []*entities.PermissionUser, err error) {
	args := x.Called(userId)

	if n, ok := args.Get(0).([]*entities.PermissionUser); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *PermissionUserRepositoryMock) DeleteByPermissionIds(ctx context.Context, ids []string) (err error) {
	args := x.Called(ids)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
