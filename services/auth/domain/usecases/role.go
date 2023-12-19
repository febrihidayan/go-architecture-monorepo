package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type RoleUsecase interface {
	Create(ctx context.Context, payload entities.RoleDto) (*entities.Role, *exceptions.CustomError)
	Find(ctx context.Context, id string) (*entities.Role, *exceptions.CustomError)
	GetAll(ctx context.Context, params entities.RoleQueryParams) (*entities.RoleMeta, *exceptions.CustomError)
}
