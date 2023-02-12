package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, payload entities.User) (*entities.User, error)
}
