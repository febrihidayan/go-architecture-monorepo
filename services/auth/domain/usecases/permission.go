package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type PermissionUsecase interface {
	Create(ctx context.Context, payload entities.PermissionDto) (*entities.Permission, *exceptions.CustomError)
	Find(ctx context.Context, id string) (*entities.Permission, *exceptions.CustomError)
	GetAll(ctx context.Context, params entities.PermissionQueryParams) (*entities.PermissionMeta, *exceptions.CustomError)
	Update(ctx context.Context, payload entities.PermissionDto) (*entities.Permission, *exceptions.CustomError)
}
