package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type RoleRepository interface {
	Create(ctx context.Context, payload *entities.Role) error
	Find(ctx context.Context, id string) (*entities.Role, error)
	FindByName(ctx context.Context, name string) (*entities.Role, error)
	Update(ctx context.Context, payload *entities.Role) error
}
