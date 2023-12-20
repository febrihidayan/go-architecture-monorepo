package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type PermissionUserRepository interface {
	CreateMany(ctx context.Context, payloads []*entities.PermissionUser) error
	AllByUserId(ctx context.Context, userId string) ([]*entities.PermissionUser, error)
	Delete(ctx context.Context, payload *entities.PermissionUser) error
}
