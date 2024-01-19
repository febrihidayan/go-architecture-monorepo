package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type PermissionRoleRepository interface {
	CreateMany(ctx context.Context, payloads []*entities.PermissionRole) error
	AllByRoleId(ctx context.Context, roleId string) ([]*entities.PermissionRole, error)
	Delete(ctx context.Context, payload *entities.PermissionRole) error
	DeleteByPermissionIds(ctx context.Context, ids []string) error
	DeleteByRoleId(ctx context.Context, roleId string) error
}
