package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
)

type AuthRepository interface {
	CreateOrUpdate(ctx context.Context, payload *entities.Auth) error
}
