package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type AclUsecase interface {
	AccessUserLogin(ctx context.Context, userId string) (*entities.AclMeta, *exceptions.CustomError)
	GetAllRole(ctx context.Context) ([]*entities.Role, *exceptions.CustomError)
	GetAllPermission(ctx context.Context) ([]*entities.Permission, *exceptions.CustomError)
	GetAllPermissionByRole(ctx context.Context, roleId string) ([]*entities.Permission, *exceptions.CustomError)
	GetAllUser(ctx context.Context, userId string) (*entities.AclMeta, *exceptions.CustomError)
	UpdateUser(ctx context.Context, payload entities.AclUserDto) *exceptions.CustomError
	UpdatePermissionByRole(ctx context.Context, payload entities.AclPermissionDto) *exceptions.CustomError
}
