package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type PermissionRepository interface {
	Create(ctx context.Context, payload *entities.Permission) error
	Find(ctx context.Context, id string) (*entities.Permission, error)
	FindByName(ctx context.Context, name string) (*entities.Permission, error)
	GetAll(ctx context.Context, params *entities.PermissionQueryParams) ([]*entities.Permission, int, error)
	Update(ctx context.Context, payload *entities.Permission) error
}
