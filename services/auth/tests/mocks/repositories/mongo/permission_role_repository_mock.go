package mongo_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

type PermissionRoleRepositoryMock struct {
	mock.Mock
}

func (x *PermissionRoleRepositoryMock) CreateMany(ctx context.Context, payload []*entities.PermissionRole) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *PermissionRoleRepositoryMock) AllByRoleId(ctx context.Context, roleId string) (result []*entities.PermissionRole, err error) {
	args := x.Called(roleId)

	if n, ok := args.Get(0).([]*entities.PermissionRole); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *PermissionRoleRepositoryMock) Delete(ctx context.Context, payload *entities.PermissionRole) (err error) {
	args := x.Called(payload.RoleId, payload.PermissionId)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *PermissionRoleRepositoryMock) DeleteByPermissionIds(ctx context.Context, ids []string) (err error) {
	args := x.Called(ids)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
