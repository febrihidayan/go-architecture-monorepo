package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type RoleUserRepository interface {
	CreateMany(ctx context.Context, payloads []*entities.RoleUser) error
	AllByUserId(ctx context.Context, userId string) ([]*entities.RoleUser, error)
	Delete(ctx context.Context, payload *entities.RoleUser) error
	DeleteByUserId(ctx context.Context, userId string) error
	DeleteByRoleIds(ctx context.Context, ids []string) error
}
